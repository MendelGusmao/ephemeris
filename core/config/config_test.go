package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

type Config struct{}

const (
	filename = "/tmp/loader-test.json"
)

func TestLoad(t *testing.T) {
	defer os.Remove(filename)

	var config Config
	err := Load(filename, &config)

	if _, ok := err.(*os.PathError); !ok {
		t.Fatal("Expecting os.PathError")
	}

	if err := ioutil.WriteFile(filename, []byte("malformed json"), 0600); err != nil {
		t.Error(err)
	}

	err = Load(filename, &config)

	if _, ok := err.(*json.SyntaxError); !ok {
		t.Fatal("Expecting json.SyntaxError")
	}

	if err := ioutil.WriteFile(filename, []byte("{}"), 0600); err != nil {
		t.Error(err)
	}

	err = Load(filename, &config)

	if err != nil {
		t.Fatal("Expecting JSON load, got", err)
	}
}
