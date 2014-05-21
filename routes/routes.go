package routes

import (
	"github.com/go-martini/martini"
)

type registrator func(m *martini.ClassicMartini)

var registrators = make([]registrator, 0)

func Register(registrator) {
	registrators = append(registrators)
}

func Apply(m *martini.ClassicMartini) {
	for _, registrator := range registrators {
		registrator(m)
	}
}
