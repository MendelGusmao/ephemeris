package main

import (
	"events/config"
	_ "events/handlers"
	"events/lib/gorm"
	"events/lib/martini"
	"events/lib/middleware/render"
	// "events/lib/middleware/sessions"
	"events/middleware"
	"events/models"
	"events/routes"
	"fmt"
	"os"
)

func main() {
	readConfiguration()
	buildDatabase()

	m := martini.Classic()

	// store := sessions.NewCookieStore([]byte(config.Events.Session.Secret))
	databaseOptions := middleware.DatabaseOptions{
		URL: config.Events.Database.URL,
	}

	// m.Use(sessions.Sessions(config.Events.Session.Name, store))
	m.Use(render.Renderer())
	m.Use(middleware.Database(databaseOptions))

	m.Group(config.Events.APIRoot, func(r martini.Router) {
		routes.Apply(r)
	})

	m.Run()
}

func readConfiguration() {
	var configFilename string
	env := os.Getenv("EVENTS_CONFIG_FILE")

	if len(env) > 0 {
		fmt.Println("Using EVENTS_CONFIG_FILE environment variable")
		configFilename = env
	}

	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}

	if len(configFilename) == 0 {
		fmt.Printf("Usage: %s <configuration file> OR set EVENTS_CONFIG_FILE environment variable\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("Using %s as configuration file\n", configFilename)

	if err := config.Load(configFilename); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}
}

func buildDatabase() {
	db, err := gorm.Open("postgres", config.Events.Database.URL)

	if err != nil {
		panic(err)
	}

	for _, model := range models.Models {
		db.AutoMigrate(model)
	}
}
