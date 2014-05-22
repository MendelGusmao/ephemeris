package protocol

import (
	"events/lib/middleware/binding"
	"net/http"
)

type EventRequest struct {
	Name                  string `json:"name"`
	Place                 string `json:"place"`
	Description           string `json:"description"`
	URL                   string `json:"url"`
	LogoURL               string `json:"logoURL"`
	Beginning             string `json:"beginning"`
	End                   string `json:"end"`
	RegistrationBeginning string `json:"registrationBeginning"`
	RegistrationEnd       string `json:"registrationEnd"`
	Visibility            string `json:"visibility"`
	Status                string `json:"status"`
}

type EventResponse struct {
	Id                    string `json:"name"`
	Name                  string `json:"name"`
	Place                 string `json:"place"`
	Description           string `json:"description"`
	URL                   string `json:"url"`
	LogoURL               string `json:"logoURL"`
	Beginning             string `json:"beginning"`
	End                   string `json:"end"`
	RegistrationBeginning string `json:"registrationBeginning"`
	RegistrationEnd       string `json:"registrationEnd"`
	Visibility            string `json:"visibility"`
	Status                string `json:"status"`
	CreatedAt             string `json:"createdAt"`
	UpdatedAt             string `json:"updatedAt"`
}

func (e *EventRequest) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	return errors
}
