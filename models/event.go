package models

import (
	"time"
)

const (
	EventVisibilityPublic  = "public"
	EventVisibilityPrivate = "private"
	EventStatusOpen        = "open"
	EventStatusOnHold      = "on hold"
	EventStatusCancelled   = "cancelled"
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
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func init() {
	register(Event{})
}
