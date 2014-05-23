package handlers

import (
	"events/lib/gorm"
	"events/lib/martini"
	"events/lib/middleware/binding"
	"events/lib/middleware/render"
	"events/models"
	"events/protocol"
	"events/protocol/transcoders"
	"events/routes"
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
	renderer render.Render,
	response http.ResponseWriter,
) {
	event, errs := transcoders.EventRequestToEvent(&eventRequest, "")

	if len(errs) > 0 {
		renderer.JSON(http.StatusBadRequest, errs)
		return
	}

	if query := database.Save(&event); query.Error != nil {
		renderer.JSON(http.StatusInternalServerError, query.Error.Error())
		return
	}

	response.Header().Add("Location", fmt.Sprintf("/events/%d", event.Id))
	response.WriteHeader(http.StatusCreated)
}

func events(
	database *gorm.DB,
	renderer render.Render,
	response http.ResponseWriter,
) {
	events := make([]models.Event, 0)

	if query := database.Find(&events); query.Error != nil {
		renderer.JSON(http.StatusInternalServerError, query.Error.Error())
		return
	}

	if len(events) == 0 {
		response.WriteHeader(http.StatusNoContent)
		return
	}

	responseEvents := make([]*protocol.EventResponse, len(events))

	for index, event := range events {
		responseEvents[index] = transcoders.EventToEventResponse(&event)
	}

	renderer.JSON(http.StatusOK, responseEvents)
}

func event(
	database *gorm.DB,
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

		renderer.JSON(http.StatusInternalServerError, query.Error.Error())
		return
	}

	renderer.JSON(http.StatusOK, transcoders.EventToEventResponse(&event))
}

func updateEvent(
	database *gorm.DB,
	eventRequest protocol.EventRequest,
	params martini.Params,
	renderer render.Render,
	response http.ResponseWriter,
) {
	if query := database.Where("id = ?", params["id"]); query.Error != nil {
		if query.Error == gorm.RecordNotFound {
			response.WriteHeader(http.StatusNotFound)
			return
		}

		renderer.JSON(http.StatusInternalServerError, query.Error.Error())
		return
	}

	event, errs := transcoders.EventRequestToEvent(&eventRequest, params["id"])

	if len(errs) > 0 {
		renderer.JSON(http.StatusBadRequest, errs)
		return
	}

	if query := database.Save(&event); query.Error != nil {
		renderer.JSON(http.StatusInternalServerError, query.Error.Error())
		return
	}

	response.Header().Add("Location", fmt.Sprintf("/events/%d", event.Id))
	response.WriteHeader(http.StatusCreated)
}

func deleteEvent(
	database *gorm.DB,
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

		renderer.JSON(http.StatusInternalServerError, query.Error.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
}
