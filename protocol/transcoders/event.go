package transcoders

import (
	"events/models"
	"events/protocol"
	"strconv"
	"time"
)

func EventRequestToEvent(from *protocol.EventRequest, id string) (*models.Event, []error) {
	errs := make([]error, 0)

	registrationBeginning, err := time.Parse(time.RFC3339, from.RegistrationBeginning)

	if err != nil {
		errs = append(errs, err)
	}

	registrationEnd, err := time.Parse(time.RFC3339, from.RegistrationEnd)

	if err != nil {
		errs = append(errs, err)
	}

	beginning, err := time.Parse(time.RFC3339, from.Beginning)

	if err != nil {
		errs = append(errs, err)
	}

	end, err := time.Parse(time.RFC3339, from.End)

	if err != nil {
		errs = append(errs, err)
	}

	to := &models.Event{
		Name:                  from.Name,
		Place:                 from.Place,
		Description:           from.Description,
		URL:                   from.URL,
		LogoURL:               from.LogoURL,
		Beginning:             beginning,
		End:                   end,
		Visibility:            models.EventVisibility(from.Visibility),
		Status:                models.EventStatus(from.Status),
		RegistrationBeginning: registrationBeginning,
		RegistrationEnd:       registrationEnd,
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

func EventToEventResponse(from *models.Event) *protocol.EventResponse {
	return &protocol.EventResponse{
		Id:                    strconv.Itoa(from.Id),
		Name:                  from.Name,
		Place:                 from.Place,
		Description:           from.Description,
		URL:                   from.URL,
		LogoURL:               from.LogoURL,
		Beginning:             from.Beginning.Format(time.RFC3339),
		End:                   from.End.Format(time.RFC3339),
		Visibility:            string(from.Visibility),
		Status:                string(from.Status),
		CreatedAt:             from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:             from.UpdatedAt.Format(time.RFC3339),
		RegistrationBeginning: from.RegistrationBeginning.Format(time.RFC3339),
		RegistrationEnd:       from.RegistrationEnd.Format(time.RFC3339),
	}
}
