package middleware

import (
	"database/sql"
	"ephemeris/core"
	"ephemeris/core/models"
	"log/syslog"
	"net/http"

	"github.com/MendelGusmao/gorm"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func Authorize(role models.UserRole) martini.Handler {
	return func(
		context martini.Context,
		database *gorm.DB,
		logger core.Logger,
		renderer render.Render,
		session sessions.Session,
	) {
		id := session.Get("user.id")

		if id == nil {
			logger.Log(syslog.LOG_DEBUG, "Not allowed")
			renderer.Status(http.StatusForbidden)
			return
		}

		user := models.User{Id: id.(int)}

		if query := database.Find(&user); query.Error != nil {
			if query.Error == gorm.RecordNotFound || query.Error == sql.ErrNoRows {
				renderer.Status(http.StatusForbidden)
				return
			}

			logger.Log(syslog.LOG_ERR, query.Error)
			renderer.Status(http.StatusInternalServerError)
			return
		}

		if user.Role < role {
			logger.Logf(syslog.LOG_WARNING, "User '%s' has no privileges for this action", user.Username)
			renderer.Status(http.StatusForbidden)
			return
		}

		logger.Logf(syslog.LOG_DEBUG, "Loading user '%s'", user.Username)
		context.Map(&user)
	}
}
