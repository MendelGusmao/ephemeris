package handlers

import (
	"ephemeris/lib/gorm"
	"ephemeris/lib/martini"
	"ephemeris/lib/middleware/binding"
	"ephemeris/lib/middleware/render"
	"ephemeris/middleware"
	"ephemeris/models"
	"ephemeris/protocol"
	"ephemeris/protocol/transcoders"
	"ephemeris/routes"
	"fmt"
	"net/http"
)

func init() {
	routes.Register(func(r martini.Router) {
		r.Get("/events", events)
		r.Post("/events", binding.Bind(protocol.EventRequest{}), createEvent)

		r.Get("/events/:id", event)
		r.Put("/events/:id", binding.Bind(protocol.EventRequest{}), updateEvent)
		r.Delete("/events/:id", deleteEvent)
	})
}

func createEvent(
	database *gorm.DB,
	eventRequest protocol.EventRequest,
	logger *middleware.ApplicationLogger,
	renderer render.Render,
	response http.ResponseWriter,
) {
	event, errs := transcoders.EventRequestToEvent(&eventRequest, "")

	if len(errs) > 0 {
		renderer.JSON(http.StatusBadRequest, errs)
		return
	}

	if query := database.Save(&event); query.Error != nil {
		logger.Log(query.Error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Add("Location", fmt.Sprintf("/events/%d", event.Id))
	response.WriteHeader(http.StatusCreated)
}

func events(
	database *gorm.DB,
	logger *middleware.ApplicationLogger,
	renderer render.Render,
	response http.ResponseWriter,
) {
	events := make([]models.Event, 0)

	if query := database.Find(&events); query.Error != nil {
		logger.Log(query.Error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(events) == 0 {
		response.WriteHeader(http.StatusNoContent)
		return
	}

	renderer.JSON(http.StatusOK, events)
}

func event(
	database *gorm.DB,
	logger *middleware.ApplicationLogger,
	params martini.Params,
	renderer render.Render,
	response http.ResponseWriter,
) {
	event := models.Event{}
	query := database.Where("id = ?", params["id"]).First(&event)

	if query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			response.WriteHeader(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderer.JSON(http.StatusOK, event)
}

func updateEvent(
	database *gorm.DB,
	eventRequest protocol.EventRequest,
	logger *middleware.ApplicationLogger,
	params martini.Params,
	renderer render.Render,
	response http.ResponseWriter,
) {
	if query := database.Where("id = ?", params["id"]); query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			response.WriteHeader(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, errs := transcoders.EventRequestToEvent(&eventRequest, params["id"])

	if len(errs) > 0 {
		renderer.JSON(http.StatusBadRequest, errs)
		return
	}

	if query := database.Save(&event); query.Error != nil {
		logger.Log(query.Error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}

func deleteEvent(
	database *gorm.DB,
	logger *middleware.ApplicationLogger,
	params martini.Params,
	renderer render.Render,
	response http.ResponseWriter,
) {
	query := database.Where("id = ?", params["id"]).Delete(&models.Event{})

	if query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			response.WriteHeader(http.StatusNotFound)
			return
		}

		logger.Log(query.Error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}
