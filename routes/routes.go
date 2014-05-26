package routes

import (
	"ephemeris/lib/martini"
)

type registrator func(r martini.Router)

var registrators = make([]registrator, 0)

func Register(r registrator) {
	registrators = append(registrators, r)
}

func Apply(r martini.Router) {
	for _, registrator := range registrators {
		registrator(r)
	}
}
