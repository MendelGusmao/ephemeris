package main

import (
	"ephemeris/core/config"
	_ "ephemeris/core/handlers"
	"ephemeris/core/middleware"
	"ephemeris/core/models"
	"ephemeris/core/routes"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"os"
)

func main() {
	readConfiguration()
	buildDatabase()

	m := martini.Classic()

	store := sessions.NewCookieStore([]byte(config.Ephemeris.Session.Secret))
	databaseOptions := middleware.DatabaseOptions{
		URL:                config.Ephemeris.Database.URL,
		MaxIdleConnections: config.Ephemeris.Database.MaxIdleConnections,
	}

	m.Use(sessions.Sessions(config.Ephemeris.Session.Name, store))
	m.Use(render.Renderer())
	m.Use(middleware.Database(databaseOptions))
	m.Use(middleware.Logger())

	if os.Getenv("DEV_RUNNER") == "1" {
		m.Use(middleware.Fresh)
	}

	m.Group(config.Ephemeris.APIRoot, func(r martini.Router) {
		routes.Apply(r)
	})

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

	if err := config.Load(configFilename); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}
}

func buildDatabase() {
	db, err := gorm.Open("postgres", config.Ephemeris.Database.URL)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	for _, err := range models.BuildDatabase(db) {
		fmt.Println("Error building database:", err)
	}
}
