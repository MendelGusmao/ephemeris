package models

import (
	"time"
)

const (
	EventVisibilityPublic  EventVisibility = "public"
	EventVisibilityPrivate EventVisibility = "private"
	EventStatusOpen        EventStatus     = "open"
	EventStatusOnHold      EventStatus     = "on hold"
	EventStatusCancelled   EventStatus     = "cancelled"
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
	Visibility            EventVisibility
	Status                EventStatus
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type EventVisibility string

type EventStatus string

func init() {
	register(Event{})
}
