package handlers

import (
	"database/sql"
	"ephemeris/core"
	"ephemeris/core/middleware"
	"ephemeris/core/models"
	"ephemeris/core/representers"
	"ephemeris/core/routes"
	"log/syslog"
	"net/http"

	"github.com/MendelGusmao/gorm"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func init() {
	routes.Register(func(r martini.Router) {
		r.Get("/session", session)
		r.Post("/session", binding.Bind(representers.UserCredentials{}), newSession)
		r.Delete("/session", middleware.Authorize(models.UserRoleNone), destroySession)
	})
}

func session(
	renderer render.Render,
	session sessions.Session,
) {
	if session.Get("user.id") == nil {
		renderer.Status(http.StatusForbidden)
		return
	}

	renderer.Status(http.StatusNoContent)
}

func newSession(
	database *gorm.DB,
	logger core.Logger,
	renderer render.Render,
	session sessions.Session,
	credentials representers.UserCredentials,
) {
	if session.Get("user.id") != nil {
		renderer.Status(http.StatusNoContent)
		return
	}

	user := models.User{}

	if query := database.Where(&models.User{
		Username: credentials.Username,
		Password: *credentials.Password,
	}).Find(&user); query.Error != nil {
		if query.Error == gorm.RecordNotFound || query.Error == sql.ErrNoRows {
			logger.Logf(syslog.LOG_WARNING, "Unsuccessful login from '%s'", user.Username)
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	logger.Logf(syslog.LOG_NOTICE, "'%s' has successfully logged in", user.Username)
	session.Set("user.id", user.Id)
	renderer.Status(http.StatusCreated)
}

func destroySession(
	logger core.Logger,
	renderer render.Render,
	session sessions.Session,
	user *models.User,
) {
	logger.Logf(syslog.LOG_NOTICE, "'%s' has successfully logged out", user.Username)
	session.Clear()
	renderer.Status(http.StatusNoContent)
}
