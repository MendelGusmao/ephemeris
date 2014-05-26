package protocol

import (
	"ephemeris/lib/middleware/binding"
	"ephemeris/models"
	"net/http"
	"net/url"
	"time"
)

type EventRequest struct {
	Name                  string    `json:"name" binding:"required"`
	Place                 string    `json:"place" binding:"required"`
	Description           string    `json:"description" binding:"required"`
	URL                   string    `json:"url" binding:"required"`
	LogoURL               string    `json:"logoURL" binding:"required"`
	Beginning             time.Time `json:"beginning" binding:"required"`
	End                   time.Time `json:"end" binding:"required"`
	RegistrationBeginning time.Time `json:"registrationBeginning" binding:"required"`
	RegistrationEnd       time.Time `json:"registrationEnd" binding:"required"`
	Visibility            string    `json:"visibility" binding:"required"`
	Status                string    `json:"status" binding:"required"`
}

func (event *EventRequest) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	errors = append(errors, event.validateDates(errors)...)
	errors = append(errors, event.validateURLs(errors)...)
	errors = append(errors, event.validateEnums(errors)...)

	return errors
}

func (event *EventRequest) validateDates(errors binding.Errors) binding.Errors {
	day, _ := time.ParseDuration("24h")
	limit := time.Now().Add(-day).Unix()

	if event.RegistrationBeginning.Unix() < limit ||
		event.RegistrationEnd.Unix() < limit ||
		event.Beginning.Unix() < limit ||
		event.End.Unix() < limit {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "registrationEnd", "beginning", "end"},
			Classification: "DateError",
			Message:        "None of the dates can be in the past",
		})
	}

	if event.RegistrationBeginning.Unix() > event.RegistrationEnd.Unix() {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "registrationEnd"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after registration end",
		})
	}

	if event.Beginning.Unix() > event.End.Unix() {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"beginning", "end"},
			Classification: "DateError",
			Message:        "Event beginning can't be after event end",
		})
	}

	if event.RegistrationEnd.Unix() > event.End.Unix() {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"registrationEnd", "end"},
			Classification: "DateError",
			Message:        "Registration end can't be after event end",
		})
	}

	if event.RegistrationBeginning.Unix() > event.Beginning.Unix() {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "beginning"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after event beginning",
		})
	}

	if event.RegistrationBeginning.Unix() > event.End.Unix() {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "end"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after event end",
		})
	}

	return errors
}

func (event *EventRequest) validateURLs(errors binding.Errors) binding.Errors {
	_, err := url.Parse(event.URL)

	if err != nil {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"url"},
			Classification: "URLError",
			Message:        "Invalid URL",
		})
	}

	_, err = url.Parse(event.LogoURL)

	if err != nil {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"logoURL"},
			Classification: "URLError",
			Message:        "Invalid URL",
		})
	}

	return errors
}

func (event *EventRequest) validateEnums(errors binding.Errors) binding.Errors {
	visibility := models.EventVisibility(event.Visibility)
	status := models.EventStatus(event.Status)

	if visibility != models.EventVisibilityPrivate ||
		visibility != models.EventVisibilityPublic {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"visibility"},
			Classification: "EnumError",
			Message:        "Invalid visibility",
		})
	}

	if status != models.EventStatusCancelled ||
		status != models.EventStatusOpen ||
		status != models.EventStatusOnHold {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"status"},
			Classification: "EnumError",
			Message:        "Invalid status",
		})
	}

	return errors
}
