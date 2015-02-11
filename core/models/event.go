package models

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/martini-contrib/binding"
)

const (
	EventVisibilityPublic  = "public"
	EventVisibilityPrivate = "private"
	EventStatusOpen        = "open"
	EventStatusOnHold      = "on hold"
	EventStatusCancelled   = "cancelled"
)

var (
	day, _            = time.ParseDuration("24h")
	validSchemes      = []string{"http", "https", "ftp", "data"}
	validVisibilities = []string{
		EventVisibilityPublic,
		EventVisibilityPrivate,
	}
	validStatuses = []string{
		EventStatusOpen,
		EventStatusOnHold,
		EventStatusCancelled,
	}
)

type Event struct {
	Id                    int
	Name                  string
	Place                 string
	Description           string
	URL                   string
	LogoURL               string
	Beginning             time.Time
	End                   time.Time
	RegistrationBeginning time.Time
	RegistrationEnd       time.Time
	Visibility            string
	Status                string
	User                  User
	UserId                int
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

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

	errors binding.Errors `json:"-"`
}

type EventUpdateRequest struct {
	EventRequest
}

func (event *Event) Update(from *EventRequest) {
	event.Name = from.Name
	event.Place = from.Place
	event.Description = from.Description
	event.URL = from.URL
	event.LogoURL = from.LogoURL
	event.Beginning = from.Beginning
	event.End = from.End
	event.Visibility = from.Visibility
	event.Status = from.Status
	event.RegistrationBeginning = from.RegistrationBeginning
	event.RegistrationEnd = from.RegistrationEnd
}

func (event *Event) ToResponse() EventResponse {
	return EventResponse{
		Id:                    event.Id,
		Name:                  event.Name,
		Place:                 event.Place,
		Description:           event.Description,
		URL:                   event.URL,
		LogoURL:               event.LogoURL,
		Beginning:             event.Beginning,
		End:                   event.End,
		Visibility:            event.Visibility,
		Status:                event.Status,
		RegistrationBeginning: event.RegistrationBeginning,
		RegistrationEnd:       event.RegistrationEnd,
		User:                  event.User.ToPrivateResponse(),
		CreatedAt:             event.CreatedAt,
		UpdatedAt:             event.UpdatedAt,
	}
}

func (event *EventRequest) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	event.errors = errors

	event.validateDates()
	event.validateURLs()
	event.validateEnums()

	return event.errors
}

func (event *EventUpdateRequest) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	event.errors = errors

	event.validateURLs()
	event.validateEnums()

	return event.errors
}

func (event *EventRequest) validateDates() {
	limit := time.Now().Add(-day).Unix()

	if event.RegistrationBeginning.Unix() < limit ||
		event.RegistrationEnd.Unix() < limit ||
		event.Beginning.Unix() < limit ||
		event.End.Unix() < limit {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "registrationEnd", "beginning", "end"},
			Classification: "DateError",
			Message:        "None of the dates can be in the past",
		})
	}

	if event.RegistrationBeginning.Unix() > event.RegistrationEnd.Unix() {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "registrationEnd"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after registration end",
		})
	}

	if event.Beginning.Unix() > event.End.Unix() {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"beginning", "end"},
			Classification: "DateError",
			Message:        "Event beginning can't be after event end",
		})
	}

	if event.RegistrationEnd.Unix() > event.End.Unix() {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"registrationEnd", "end"},
			Classification: "DateError",
			Message:        "Registration end can't be after event end",
		})
	}

	if event.RegistrationBeginning.Unix() > event.Beginning.Unix() {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "beginning"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after event beginning",
		})
	}

	if event.RegistrationBeginning.Unix() > event.End.Unix() {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"registrationBeginning", "end"},
			Classification: "DateError",
			Message:        "Registration beginning can't be after event end",
		})
	}
}

func (event *EventRequest) validateURLs() {
	event.validateURL(event.URL, "URL")
	event.validateURL(event.LogoURL, "LogoURL")
}

func (event *EventRequest) validateURL(value, field string) {
	uri, err := url.Parse(value)

	if err != nil {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{field},
			Classification: "URLError",
			Message:        "Invalid URL",
		})

		return
	}

	ok := false

	for _, scheme := range validSchemes {
		if scheme == uri.Scheme {
			ok = true
			break
		}
	}

	if !ok {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{field},
			Classification: "URLError",
			Message:        "Invalid URL Scheme",
		})
	}
}

func (event *EventRequest) validateEnums() {
	ok := false

	for _, visibility := range validVisibilities {
		if event.Visibility == visibility {
			ok = true
			break
		}
	}

	if !ok {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"visibility"},
			Classification: "EnumError",
			Message:        fmt.Sprintf("Invalid visibility '%s'", event.Visibility),
		})
	}

	ok = false

	for _, status := range validStatuses {
		if event.Status == status {
			ok = true
			break
		}
	}

	if !ok {
		event.errors = append(event.errors, binding.Error{
			FieldNames:     []string{"status"},
			Classification: "EnumError",
			Message:        fmt.Sprintf("Invalid status '%s'", event.Status),
		})
	}
}

func init() {
	register(&Event{})
}
