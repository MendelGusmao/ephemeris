package models

import (
	"time"
)

type User struct {
	Id            int
	Username      string
	Password      string
	Administrator bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func init() {
	register(User{})
}
