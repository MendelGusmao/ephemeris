package transcoders

import (
	"ephemeris/models"
	"ephemeris/protocol"
)

func EventRequestToEvent(from *protocol.EventRequest, to *models.Event) {
	to.Name = from.Name
	to.Place = from.Place
	to.Description = from.Description
	to.URL = from.URL
	to.LogoURL = from.LogoURL
	to.Beginning = from.Beginning
	to.End = from.End
	to.Visibility = from.Visibility
	to.Status = from.Status
	to.RegistrationBeginning = from.RegistrationBeginning
	to.RegistrationEnd = from.RegistrationEnd
}
