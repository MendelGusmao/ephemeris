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
	routes.Register(func(m *martini.ClassicMartini) {
		m.Get("/events", events)
		m.Post("/events", binding.Bind(protocol.EventRequest{}), createEvent)

		m.Get("/event/:id", event)
		m.Patch("/event/:id", binding.Bind(protocol.EventRequest{}), updateEvent)
		m.Delete("/event/:id", deleteEvent)
	})
}

func createEvent(
	database *gorm.DB,
	eventRequest protocol.EventRequest,
	renderer render.Render,
	response http.ResponseWriter,
) {
	event, errs := transcoders.EventRequestToEvent(&eventRequest)

	if len(errs) > 0 {
		renderer.JSON(http.StatusInternalServerError, errs)
		return
	}

	err := database.Save(&event)

	if err != nil {
		renderer.JSON(http.StatusInternalServerError, err)
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
	database.Find(&events)

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

	renderer.JSON(200, transcoders.EventToEventResponse(&event))
}

func updateEvent(
	database *gorm.DB,
	renderer render.Render,
) {

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
