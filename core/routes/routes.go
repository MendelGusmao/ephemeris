package routes

import (
	"github.com/go-martini/martini"
	"sync"
)

var (
	registerLock sync.RWMutex
)

type registrator func(r martini.Router)

var registrators = make([]registrator, 0)

func Register(r registrator) {
	registerLock.Lock()
	defer registerLock.Unlock()

	registrators = append(registrators, r)
}

func Apply(r martini.Router) {
	registerLock.Lock()
	defer registerLock.Unlock()

	for _, registrator := range registrators {
		registrator(r)
	}
}
