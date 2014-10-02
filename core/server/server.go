package server

import (
	"ephemeris/core/config"
	"ephemeris/core/middleware"
	"ephemeris/core/routes"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func Configure(ephemeris config.EphemerisConfig, m *martini.ClassicMartini) {
	store := sessions.NewCookieStore([]byte(ephemeris.Session.Secret))
	dbOptions := middleware.DBOptions{
		Driver:             ephemeris.Database.Driver,
		URL:                ephemeris.Database.URL,
		MaxIdleConnections: ephemeris.Database.MaxIdleConnections,
	}

	m.Use(sessions.Sessions(ephemeris.Session.Name, store))
	m.Use(render.Renderer())
	m.Use(middleware.Database(dbOptions))
	m.Use(middleware.Logger())

	if os.Getenv("DEV_RUNNER") == "1" {
		m.Use(middleware.Fresh)
	}

	m.Group(ephemeris.APIRoot, func(r martini.Router) {
		routes.Apply(r)
	})
}
