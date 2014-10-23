package representers

import (
	"ephemeris/core/models"
	"fmt"
	"net/http"

	"github.com/martini-contrib/binding"
)

type UserResponse struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     int    `json:"role,omitempty"`
}

type UserRequest struct {
	Username string  `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     int     `json:"roles,omitempty"`
}

type UserCredentials struct {
	UserRequest
}

func (user *UserRequest) Validate(errors binding.Errors, request *http.Request) binding.Errors {
	switch models.UserRole(user.Role) {
	case models.UserRoleRegular, models.UserRoleAccrediting,
		models.UserRoleManager, models.UserRoleAdministrator:
	default:
		errors = append(errors, binding.Error{
			FieldNames:     []string{"roles"},
			Classification: "RolesError",
			Message:        fmt.Sprintf("Invalid role: '%d'", user.Role),
		})
	}

	return errors
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
