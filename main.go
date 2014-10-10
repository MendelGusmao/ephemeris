package main

import (
	"ephemeris/core/config"
	_ "ephemeris/core/handlers"
	"ephemeris/core/models"
	"ephemeris/core/server"
	"fmt"
	"os"

	"github.com/MendelGusmao/gorm"
	"github.com/go-martini/martini"
)

func main() {
	readConfiguration()
	buildDatabase()

	m := martini.Classic()
	server.Configure(config.Ephemeris, m)
	m.Run()
}

func readConfiguration() {
	var configFilename string
	env := os.Getenv("EPHEMERIS_CONFIG_FILE")

	if len(env) > 0 {
		fmt.Println("Using EPHEMERIS_CONFIG_FILE environment variable")
		configFilename = env
	}

	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}

	if len(configFilename) == 0 {
		fmt.Printf("Usage: %s <configuration file> OR set EPHEMERIS_CONFIG_FILE environment variable\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("Using %s as configuration file\n", configFilename)

	if err := config.Load(configFilename, &config.Ephemeris); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}
}

func buildDatabase() {
	db, err := gorm.Open(config.Ephemeris.Database.Driver, config.Ephemeris.Database.URL)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	for _, err := range models.BuildDatabase(db) {
		fmt.Println("Error building database:", err)
	}
}
