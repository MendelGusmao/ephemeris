package handlers

import (
	"ephemeris/core/middleware"
	"ephemeris/core/models"
	"ephemeris/core/representers"
	"ephemeris/core/representers/transcoders"
	"ephemeris/core/routes"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func init() {
	routes.Register(func(r martini.Router) {
		r.Get("/session", session)
		r.Post("/session", binding.Bind(representers.UserRequest{}), newSession)
		r.Delete("/session", middleware.Authorize(), destroySession)
	})
}

func session(
	renderer render.Render,
	session sessions.Session,
) {
	if session.Get("user.id") != nil {
		renderer.Status(http.StatusNoContent)
		return
	}

	renderer.Status(http.StatusForbidden)
}

func newSession(
	database *gorm.DB,
	logger *middleware.ApplicationLogger,
	renderer render.Render,
	session sessions.Session,
	userRequest representers.UserRequest,
) {
	if session.Get("user.id") != nil {
		renderer.Status(http.StatusNoContent)
		return
	}

	user := models.User{}
	transcoders.UserFromRequest(&userRequest, &user)

	query := database.Where(
		&models.User{
			Username: user.Username,
			Password: user.Password,
		},
	).First(&user)

	if query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			logger.Logf("Unsuccessful login from '%s'", user.Username)
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	logger.Logf("'%s' has successfully logged in", user.Username)
	session.Set("user.id", user.Id)
	renderer.Status(http.StatusCreated)
}

func destroySession(
	logger *middleware.ApplicationLogger,
	renderer render.Render,
	session sessions.Session,
	user *models.User,
) {
	logger.Logf("'%s' has successfully logged out", user.Username)
	session.Clear()
	renderer.Status(http.StatusNoContent)
}
