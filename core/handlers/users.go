package handlers

import (
	"ephemeris/core/middleware"
	"ephemeris/core/models"
	"ephemeris/core/representers"
	"ephemeris/core/representers/transcoders"
	"ephemeris/core/routes"
	"fmt"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func init() {
	routes.Register(func(r martini.Router) {
		r.Get("/users", users)
		r.Post("/users", binding.Bind(representers.UserRequest{}), createUser)

		r.Get("/users/:id", user)
		r.Put("/users/:id", binding.Bind(representers.UserRequest{}), updateUser)
		r.Delete("/users/:id", deleteUser)
	})
}

func createUser(
	database *gorm.DB,
	userRequest representers.UserRequest,
	logger *middleware.AppLogger,
	renderer render.Render,
) {
	user := models.User{}
	transcoders.UserFromRequest(&userRequest, &user)

	if query := database.Save(&user); query.Error != nil {
		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Header().Add("Location", fmt.Sprintf("/users/%d", user.Id))
	renderer.Status(http.StatusCreated)
}

func users(
	database *gorm.DB,
	logger *middleware.AppLogger,
	renderer render.Render,
) {
	users := make([]models.User, 0)
	lastModified := time.Time{}

	if query := database.Find(&users); query.Error != nil {
		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	representedUsers := make([]representers.UserResponse, len(users))

	for index, user := range users {
		if user.UpdatedAt.Unix() > lastModified.Unix() {
			lastModified = user.UpdatedAt
		}

		representedUsers[index] = transcoders.UserToResponse(&user)
	}

	renderer.Header().Add("Last-Modified", lastModified.UTC().Format(time.RFC1123))
	renderer.JSON(http.StatusOK, representedUsers)
}

func user(
	database *gorm.DB,
	logger *middleware.AppLogger,
	params martini.Params,
	renderer render.Render,
) {
	user := models.User{}
	query := database.Where("id = ?", params["id"]).First(&user)

	if query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Header().Add("Last-Modified", user.CreatedAt.UTC().Format(time.RFC1123))
	renderer.JSON(http.StatusOK, transcoders.UserToResponse(&user))
}

func updateUser(
	database *gorm.DB,
	userRequest representers.UserRequest,
	logger *middleware.AppLogger,
	params martini.Params,
	renderer render.Render,
) {
	user := models.User{}

	if query := database.Where("id = ?", params["id"]).Find(&user); query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	transcoders.UserFromRequest(&userRequest, &user)

	if query := database.Save(user); query.Error != nil {
		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Header().Add("Location", fmt.Sprintf("/users/%d", user.Id))
	renderer.Status(http.StatusOK)
}

func deleteUser(
	database *gorm.DB,
	logger *middleware.AppLogger,
	params martini.Params,
	renderer render.Render,
) {
	user := models.User{}

	if query := database.Where("id = ?", params["id"]).Find(&user); query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			renderer.Status(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	if query := database.Where("id = ?", params["id"]).Delete(&user); query.Error != nil {
		logger.Log(query.Error.Error())
		renderer.Status(http.StatusInternalServerError)
		return
	}

	renderer.Status(http.StatusOK)
}
