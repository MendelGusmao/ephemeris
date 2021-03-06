package handlers

import (
	"database/sql"
	"ephemeris/core"
	"ephemeris/core/middleware"
	"ephemeris/core/models"
	"ephemeris/core/routes"
	"fmt"
	"log/syslog"
	"net/http"
	"strconv"
	"time"

	"github.com/MendelGusmao/gorm"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func init() {
	routes.Register(func(r martini.Router) {
		r.Get("/users", middleware.Authorize(models.UserRoleAdministrator), users)
		r.Post("/users", middleware.Authorize(models.UserRoleAdministrator),
			binding.Bind(models.UserRequest{}), createUser)

		r.Get("/users/:id", middleware.Authorize(models.UserRoleRegular), user)
		r.Put("/users/:id", middleware.Authorize(models.UserRoleAdministrator),
			binding.Bind(models.UserRequest{}), updateUser)
		r.Delete("/users/:id", middleware.Authorize(models.UserRoleAdministrator), deleteUser)
	})
}

func createUser(
	database *gorm.DB,
	userRequest models.UserRequest,
	logger core.Logger,
	renderer render.Render,
) {
	user := models.User{}
	user.Update(&userRequest)

	if query := database.Save(&user); query.Error != nil {
		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Header().Add("Location", fmt.Sprintf("/users/%d", user.Id))
	renderer.Status(http.StatusCreated)
}

func users(
	database *gorm.DB,
	logger core.Logger,
	renderer render.Render,
) {
	users := make([]models.User, 0)
	lastModified := time.Time{}

	if query := database.Find(&users); query.Error != nil {
		// TODO gorm doesn't return gorm.RecordNotFound when using testdb as driver
		if query.Error == gorm.RecordNotFound || query.Error == sql.ErrNoRows {
			renderer.Status(http.StatusNoContent)
			return
		}

		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	representedUsers := make([]models.UserResponse, len(users))

	for index, user := range users {
		if user.UpdatedAt.Unix() > lastModified.Unix() {
			lastModified = user.UpdatedAt
		}

		representedUsers[index] = user.ToResponse()
	}

	renderer.Header().Add("Last-Modified", lastModified.UTC().Format(time.RFC1123))
	renderer.JSON(http.StatusOK, representedUsers)
}

func user(
	database *gorm.DB,
	logger core.Logger,
	params martini.Params,
	renderer render.Render,
) {
	id, _ := strconv.Atoi(params["id"])
	user := models.User{Id: id}
	query := database.Find(&user)

	if query.Error != nil {
		// TODO gorm doesn't return gorm.RecordNotFound when using testdb as driver
		if query.Error == gorm.RecordNotFound || query.Error == sql.ErrNoRows {
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Header().Add("Last-Modified", user.CreatedAt.UTC().Format(time.RFC1123))
	renderer.JSON(http.StatusOK, user.ToResponse())
}

func updateUser(
	database *gorm.DB,
	userRequest models.UserRequest,
	logger core.Logger,
	params martini.Params,
	renderer render.Render,
) {
	id, _ := strconv.Atoi(params["id"])
	user := models.User{Id: id}

	if query := database.Find(&user); query.Error != nil {
		if query.Error == gorm.RecordNotFound || query.Error == sql.ErrNoRows {
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	user.Update(&userRequest)

	if query := database.Save(user); query.Error != nil {
		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Header().Add("Location", fmt.Sprintf("/users/%d", user.Id))
	renderer.Status(http.StatusNoContent)
}

func deleteUser(
	database *gorm.DB,
	logger core.Logger,
	params martini.Params,
	renderer render.Render,
) {
	id, _ := strconv.Atoi(params["id"])
	user := models.User{Id: id}

	if query := database.Find(&user); query.Error != nil {
		// TODO gorm doesn't return gorm.RecordNotFound when using testdb as driver
		if query.Error == gorm.RecordNotFound || query.Error == sql.ErrNoRows {
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	if query := database.Delete(&user); query.Error != nil {
		logger.Log(syslog.LOG_ERR, query.Error)
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Status(http.StatusNoContent)
}
