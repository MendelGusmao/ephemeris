package middleware

import (
	"ephemeris/core/models"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Authorize() martini.Handler {
	return func(
		context martini.Context,
		database *gorm.DB,
		logger *AppLogger,
		renderer render.Render,
		session sessions.Session,
	) {
		id := session.Get("user.id")

		if id == nil {
			logger.Log("Not allowed")
			renderer.Status(http.StatusForbidden)
			return
		}

		user, err := loadUser(database, id)

		if err != nil {
			if err == gorm.RecordNotFound {
				renderer.Status(http.StatusForbidden)
				return
			}

			logger.Log(err.Error())
			renderer.Status(http.StatusInternalServerError)
			return
		}

		logger.Logf("Loading user '%s'", user.Username)
		context.Map(user)
	}
}

func loadUser(database *gorm.DB, id interface{}) (*models.User, error) {
	user := &models.User{}
	query := database.Where("id = ?", id).Find(user)

	if query.Error != nil {
		return nil, query.Error
	}

	return user, nil
}
