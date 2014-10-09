package middleware

import (
	"ephemeris/core"
	"ephemeris/core/models"
	"log/syslog"
	"net/http"

	"github.com/MendelGusmao/gorm"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func Authorize() martini.Handler {
	return func(
		context martini.Context,
		database *gorm.DB,
		logger core.Logger,
		renderer render.Render,
		session sessions.Session,
	) {
		id := session.Get("user.id")

		if id == nil {
			logger.Log(syslog.LOG_INFO, "Not allowed")
			renderer.Status(http.StatusForbidden)
			return
		}

		user, err := loadUser(database, id)

		if err != nil {
			if err == gorm.RecordNotFound {
				renderer.Status(http.StatusForbidden)
				return
			}

			logger.Log(syslog.LOG_INFO, err.Error())
			renderer.Status(http.StatusInternalServerError)
			return
		}

		logger.Logf(syslog.LOG_INFO, "Loading user '%s'", user.Username)
		context.Map(user)
	}
}

func loadUser(database *gorm.DB, id interface{}) (*models.User, error) {
	user := &models.User{}
	query := database.Where("(`id` = ?)", id).Find(user)

	if query.Error != nil {
		return nil, query.Error
	}

	return user, nil
}
