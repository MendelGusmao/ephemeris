package config

import (
	"encoding/json"
	"io/ioutil"
)

var Events eventsConfig

type eventsConfig struct {
	APIRoot  string
	Database databaseConfig
	Session  sessionConfig
}

type databaseConfig struct {
	URL string
}

type sessionConfig struct {
	Secret string
	Name   string
}

func Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &Events)
	if err != nil {
		return err
	}

	return nil
}
