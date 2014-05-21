package models

import (
	"time"
)

var (
	EventVisibilityPublic  EventVisibility = "public"
	EventVisibilityPrivate EventVisibility = "private"
	EventStatusOpen        EventStatus     = "open"
	EventStatusOnHold      EventStatus     = "on hold"
	EventStatusCancelled   EventStatus     = "cancelled"
)

type Event struct {
	Id                    string          `json:"id"`
	Name                  string          `json:"name"`
	Place                 string          `json:"place"`
	Description           string          `json:"description"`
	URL                   string          `json:"url"`
	LogoURL               string          `json:"logoURL"`
	Beginning             time.Time       `json:"beginning"`
	End                   time.Time       `json:"end"`
	RegistrationBeginning time.Time       `json:"registrationBeginning"`
	RegistrationEnd       time.Time       `json:"registrationEnd"`
	Visibility            EventVisibility `json:"visibility"`
	Status                EventStatus     `json:"status"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
}

type EventVisibility string

type EventStatus string
