package config

import (
	"encoding/json"
	"io/ioutil"
)

var Ephemeris EphemerisConfig

type EphemerisConfig struct {
	APIRoot  string
	Database DatabaseConfig
	Session  SessionConfig
}

type DatabaseConfig struct {
	Driver             string
	URL                string
	MaxIdleConnections int
}

type SessionConfig struct {
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
