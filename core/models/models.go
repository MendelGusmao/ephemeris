package models

import (
	"sync"
)

var (
	Models       = make([]interface{}, 0)
	registerLock sync.RWMutex
)

func register(model interface{}) {
	registerLock.Lock()
	defer registerLock.Unlock()

	Models = append(Models, model)
}
