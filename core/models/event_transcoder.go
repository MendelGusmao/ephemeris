package models

func EventFromRequest(from *EventRequest, to *Event) {
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

func EventToResponse(from *Event) EventResponse {
	return EventResponse{
		Id:                    from.Id,
		Name:                  from.Name,
		Place:                 from.Place,
		Description:           from.Description,
		URL:                   from.URL,
		LogoURL:               from.LogoURL,
		Beginning:             from.Beginning,
		End:                   from.End,
		Visibility:            from.Visibility,
		Status:                from.Status,
		RegistrationBeginning: from.RegistrationBeginning,
		RegistrationEnd:       from.RegistrationEnd,
		User:                  UserToResponsePrivate(&from.User),
		CreatedAt:             from.CreatedAt,
		UpdatedAt:             from.UpdatedAt,
	}
}
