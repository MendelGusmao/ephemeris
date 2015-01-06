package models

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/martini-contrib/binding"
)

const (
	UserRoleNone          UserRole = 0
	UserRoleRegular       UserRole = 1
	UserRoleAccrediting   UserRole = 2
	UserRoleManager       UserRole = 3
	UserRoleAdministrator UserRole = 4
)

type UserRole int

type User struct {
	Id        int
	Username  string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ur *UserRole) Scan(value interface{}) error {
	if value == nil {
		*ur = UserRoleNone
		return nil
	}

	switch value.(type) {
	case int64:
		*ur = UserRole(value.(int64))
	case string:
		v := value.(string)
		role, err := strconv.Atoi(v)

		if err != nil {
			return err
		}

		*ur = UserRole(role)
	default:
		return fmt.Errorf("Error scanning user role")
	}

	return nil
}

func (ur UserRole) String() string {
	switch ur {
	case UserRoleNone:
		return "none"
	case UserRoleRegular:
		return "regular"
	case UserRoleAccrediting:
		return "accrediting"
	case UserRoleManager:
		return "manager"
	case UserRoleAdministrator:
		return "administrator"
	}

	return ""
}

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
	switch UserRole(user.Role) {
	case UserRoleRegular, UserRoleAccrediting,
		UserRoleManager, UserRoleAdministrator:
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

func init() {
	register(User{})
}
