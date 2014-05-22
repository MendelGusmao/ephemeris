package main

import (
	"events/config"
	_ "events/handlers"
	"events/middleware"
	"events/routes"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"os"
)

func main() {
	readConfiguration()

	m := martini.Classic()

	store := sessions.NewCookieStore([]byte(config.Events.Session.Secret))
	databaseOptions := middleware.DatabaseOptions{
		URL:       config.Events.Database.URL,
		Name:      config.Events.Database.Name,
		Monotonic: config.Events.Database.Monotonic,
	}

	m.Use(sessions.Sessions(config.Events.Session.Name, store))
	m.Use(render.Renderer())
	m.Use(middleware.Database(databaseOptions))

	routes.Apply(m)

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
