package models

import (
	"github.com/MendelGusmao/gorm"
)

var Models = make([]interface{}, 0)

func register(model interface{}) {
	Models = append(Models, model)
}

func BuildDatabase(db gorm.DB) []error {
	errors := make([]error, 0)

	for _, model := range Models {
		db.AutoMigrate(model)

		if db.Error != nil {
			errors = append(errors, db.Error)
		}
	}

	return errors
}
