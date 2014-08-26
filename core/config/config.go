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

func Load(filename string, config interface{}) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, config)
	if err != nil {
		return err
	}

	return nil
}
