package representers

import (
	"net/http"

	"github.com/martini-contrib/binding"
)

type UserResponse struct {
	Id            int    `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Administrator bool   `json:"administrator,omitempty"`
}

type UserRequest struct {
	Username      string  `json:"username,omitempty"`
	Password      *string `json:"password,omitempty"`
	Administrator bool    `json:"administrator,omitempty"`
}

type UserCredentials struct {
	UserRequest
}

func (credentials *UserCredentials) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	if credentials.Username == "" {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"username"},
			Classification: "StringError",
			Message:        "Username can't be empty",
		})
	}

	if credentials.Password == nil || *credentials.Password == "" {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"password"},
			Classification: "StringError",
			Message:        "Password can't be empty",
		})
	}

	return errors
}
