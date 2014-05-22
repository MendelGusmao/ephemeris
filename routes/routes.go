package routes

import (
	"events/lib/martini"
)

type registrator func(m *martini.ClassicMartini)

var registrators = make([]registrator, 0)

func Register(r registrator) {
	registrators = append(registrators, r)
}

func Apply(m *martini.ClassicMartini) {
	for _, registrator := range registrators {
		registrator(m)
	}
}
