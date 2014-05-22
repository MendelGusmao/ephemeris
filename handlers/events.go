package handlers

import (
	"events/lib/gorm"
	"events/lib/martini"
	"events/lib/middleware/binding"
	"events/lib/middleware/render"
	"events/lib/middleware/sessions"
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
		m.Post("/events", binding.Bind(protocol.EventRequest{}), events_create)

		m.Get("/event/:id", event)
		m.Patch("/event/:id", binding.Bind(protocol.EventRequest{}), event_update)
		m.Delete("/event/:id", event_delete)
	})
}

func events_create(
	eventRequest protocol.EventRequest,
	session sessions.Session,
	renderer render.Render,
	database *gorm.DB,
	response http.ResponseWriter,
) {
	event, errs := transcoders.EventRequestToEvent(&eventRequest)

	if len(errs) > 0 {
		renderer.JSON(500, errs)
		return
	}

	err := database.Save(&event)

	if err != nil {
		renderer.JSON(500, err)
		return
	}

	response.Header().Add("Location", fmt.Sprintf("/events/%d", event.Id))
	response.WriteHeader(201)
}

func events(
	response martini.ResponseWriter,
	session sessions.Session,
	renderer render.Render,
	database *gorm.DB,
) {

}

func event(
	params martini.Params,
	session sessions.Session,
	renderer render.Render,
	database *gorm.DB,
) {
	event := models.Event{}
	database.Where("id = ?", params["id"]).First(&event)
	renderer.JSON(200, transcoders.EventToEventResponse(&event))
}

func event_update(
	session sessions.Session,
	renderer render.Render,
	database *gorm.DB,
) {

}

func event_delete(
	session sessions.Session,
	renderer render.Render,
	database *gorm.DB,
) {

}
