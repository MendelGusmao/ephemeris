package models

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	models       = make([]interface{}, 0)
	registerLock sync.RWMutex
)

func register(model interface{}) {
	registerLock.Lock()
	defer registerLock.Unlock()

	models = append(models, model)
}

func BuildDatabase(db gorm.DB) []error {
	errors := make([]error, 0)

	for _, model := range models {
		db.AutoMigrate(model)

		if db.Error != nil {
			errors = append(errors, db.Error)
		}
	}

	return errors
}
