package representers

import (
	"ephemeris/core/models"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/martini-contrib/binding"
)

var (
	day, _ = time.ParseDuration("24h")
)

type EventResponse struct {
	Id                    int          `json:"id"`
	Name                  string       `json:"name"`
	Place                 string       `json:"place"`
	Description           string       `json:"description"`
	URL                   string       `json:"URL"`
	LogoURL               string       `json:"logoURL"`
	Beginning             time.Time    `json:"beginning"`
	End                   time.Time    `json:"end"`
	RegistrationBeginning time.Time    `json:"registrationBeginning"`
	RegistrationEnd       time.Time    `json:"registrationEnd"`
	Visibility            string       `json:"visibility"`
	Status                string       `json:"status"`
	User                  UserResponse `json:"user"`
	CreatedAt             time.Time    `json:"createdAt"`
	UpdatedAt             time.Time    `json:"updatedAt"`
}

type EventRequest struct {
	Name                  string    `json:"name" binding:"required"`
	Place                 string    `json:"place" binding:"required"`
	Description           string    `json:"description" binding:"required"`
	URL                   string    `json:"URL"`
	LogoURL               string    `json:"logoURL"`
	Beginning             time.Time `json:"beginning" binding:"required"`
	End                   time.Time `json:"end" binding:"required"`
	RegistrationBeginning time.Time `json:"registrationBeginning" binding:"required"`
	RegistrationEnd       time.Time `json:"registrationEnd" binding:"required"`
	Visibility            string    `json:"visibility" binding:"required"`
	Status                string    `json:"status" binding:"required"`
}

func (event *EventRequest) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	event.validateDates(&errors)
	event.validateURLs(&errors)
	event.validateEnums(&errors)

	return errors
}

func (event *EventRequest) validateDates(errors *binding.Errors) {
	limit := time.Now().Add(-day).Unix()

	if event.RegistrationBeginning.Unix() < limit ||
		event.RegistrationEnd.Unix() < limit ||
		event.Beginning.Unix() < limit ||
		event.End.Unix() < limit {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "registrationEnd", "beginning", "end"},
			Classification: "DateError",
			Message:        "None of the dates can be in the past",
		})
	}

	if event.RegistrationBeginning.Unix() > event.RegistrationEnd.Unix() {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "registrationEnd"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after registration end",
		})
	}

	if event.Beginning.Unix() > event.End.Unix() {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"beginning", "end"},
			Classification: "DateError",
			Message:        "Event beginning can't be after event end",
		})
	}

	if event.RegistrationEnd.Unix() > event.End.Unix() {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"registrationEnd", "end"},
			Classification: "DateError",
			Message:        "Registration end can't be after event end",
		})
	}

	if event.RegistrationBeginning.Unix() > event.Beginning.Unix() {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "beginning"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after event beginning",
		})
	}

	if event.RegistrationBeginning.Unix() > event.End.Unix() {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "end"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after event end",
		})
	}
}

func (event *EventRequest) validateURLs(errors *binding.Errors) {
	_, err := url.Parse(event.URL)

	if err != nil {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"url"},
			Classification: "URLError",
			Message:        "Invalid URL",
		})
	}

	_, err = url.Parse(event.LogoURL)

	if err != nil {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"logoURL"},
			Classification: "URLError",
			Message:        "Invalid URL",
		})
	}
}

func (event *EventRequest) validateEnums(errors *binding.Errors) {
	if event.Visibility != models.EventVisibilityPrivate &&
		event.Visibility != models.EventVisibilityPublic {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"visibility"},
			Classification: "EnumError",
			Message:        fmt.Sprintf("Invalid visibility '%s'", event.Visibility),
		})
	}

	if event.Status != models.EventStatusCancelled &&
		event.Status != models.EventStatusOpen &&
		event.Status != models.EventStatusOnHold {
		*errors = append(*errors, binding.Error{
			FieldNames:     []string{"status"},
			Classification: "EnumError",
			Message:        fmt.Sprintf("Invalid status '%s'", event.Status),
		})
	}
}
