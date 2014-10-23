package models

import (
	"fmt"
	"strconv"
	"time"
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

func init() {
	register(User{})
}
