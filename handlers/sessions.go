package handlers

import (
	"ephemeris/middleware"
	"ephemeris/models"
	"ephemeris/representers"
	"ephemeris/representers/transcoders"
	"ephemeris/routes"
	"fmt"
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
		r.Delete("/session", destroySession)
	})
}

func session(
	renderer render.Render,
	session sessions.Session,
) {
	if session.Get("user.id") != nil {
		renderer.Status(http.StatusOK)
		return
	}

	renderer.Status(http.StatusNotFound)
}

func newSession(
	database *gorm.DB,
	logger *middleware.ApplicationLogger,
	renderer render.Render,
	session sessions.Session,
	userRequest representers.UserRequest,
) {
	if session.Get("user.id") != nil {
		renderer.Status(http.StatusOK)
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
			logger.Log(fmt.Sprintf("Unsuccessful login from '%s'", user.Username))
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	logger.Log(fmt.Sprintf("'%s' has successfully logged in", user.Username))
	session.Set("user.id", user.Id)
	session.Set("user.name", user.Username)
	session.Set("user.administrator", user.Administrator)
	renderer.Status(http.StatusCreated)
}

func destroySession(
	logger *middleware.ApplicationLogger,
	renderer render.Render,
	session sessions.Session,
) {
	if session.Get("user.id") == nil {
		renderer.Status(http.StatusNotFound)
		return
	}

	logger.Log(fmt.Sprintf("'%s' has successfully logged out", session.Get("user.name")))
	session.Clear()
	renderer.Status(http.StatusOK)
}
