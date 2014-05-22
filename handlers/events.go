package handlers

import (
	"events/models"
	"events/protocol"
	"events/protocol/transcoders"
	"events/routes"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

const (
	eventsCollection = "events"
)

func init() {
	routes.Register(func(m *martini.ClassicMartini) {
		m.Get("/events", events_list)
		m.Post("/events", binding.Bind(protocol.EventRequest{}), events_create)

		m.Get("/event/:id", event)
		m.Patch("/event/:id", event_update)
		m.Delete("/event/:id", event_delete)
	})
}

func events_create(
	eventRequest protocol.EventRequest,
	session sessions.Session,
	renderer render.Render,
	database *mgo.Database,
	response http.ResponseWriter,
) {
	collection := database.C(eventsCollection)
	event, errs := transcoders.EventRequestToEvent(&eventRequest)

	if len(errs) > 0 {
		renderer.JSON(500, errs)
		return
	}

	err := collection.Insert(event)

	if err != nil {
		renderer.JSON(500, err)
		return
	}

	response.Header().Add("Location", "/events/"+event.Id)
	response.WriteHeader(201)
}

func events_list(
	session sessions.Session,
	renderer render.Render,
	database *mgo.Database,
) {

}

func event(
	params martini.Params,
	session sessions.Session,
	renderer render.Render,
	database *mgo.Database,
) {
	collection := database.C(eventsCollection)
	event := models.Event{}

	err := collection.Find(bson.M{"id": params["id"]}).One(&event)

	if err != nil {
		renderer.JSON(500, err)
		return
	}

	renderer.JSON(200, transcoders.EventToEventResponse(&event))
}

func event_update(
	session sessions.Session,
	renderer render.Render,
	database *mgo.Database,
) {

}

func event_delete(
	session sessions.Session,
	renderer render.Render,
	database *mgo.Database,
) {

}
