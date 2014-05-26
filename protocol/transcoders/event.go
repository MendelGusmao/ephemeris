package transcoders

import (
	"events/models"
	"events/protocol"
	"strconv"
)

func EventRequestToEvent(from *protocol.EventRequest, id string) (*models.Event, []error) {
	errs := make([]error, 0)

	to := &models.Event{
		Name:                  from.Name,
		Place:                 from.Place,
		Description:           from.Description,
		URL:                   from.URL,
		LogoURL:               from.LogoURL,
		Beginning:             from.Beginning,
		End:                   from.End,
		Visibility:            models.EventVisibility(from.Visibility),
		Status:                models.EventStatus(from.Status),
		RegistrationBeginning: from.RegistrationBeginning,
		RegistrationEnd:       from.RegistrationEnd,
	}

	if len(id) > 0 {
		intId, err := strconv.Atoi(id)

		if err != nil {
			errs = append(errs, err)
		}

		to.Id = intId
	}

	return to, errs
}
