package config

import (
	"encoding/json"
	"io/ioutil"
)

var Ephemeris ephemerisConfig

type ephemerisConfig struct {
	APIRoot  string
	Database databaseConfig
	Session  sessionConfig
}

type databaseConfig struct {
	URL                string
	MaxIdleConnections int
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

	err = json.Unmarshal(content, &Ephemeris)
	if err != nil {
		return err
	}

	return nil
}
