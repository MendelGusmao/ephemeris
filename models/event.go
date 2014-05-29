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
	Id                    int       `json:"id"`
	Name                  string    `json:"name"`
	Place                 string    `json:"place"`
	Description           string    `json:"description"`
	URL                   string    `json:"URL"`
	LogoURL               string    `json:"logoURL"`
	Beginning             time.Time `json:"beginning"`
	End                   time.Time `json:"end"`
	RegistrationBeginning time.Time `json:"registrationBeginning"`
	RegistrationEnd       time.Time `json:"registrationEnd"`
	Visibility            string    `json:"visibility"`
	Status                string    `json:"status"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

func init() {
	register(Event{})
}
