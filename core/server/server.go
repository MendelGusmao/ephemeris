package server

import (
	"ephemeris/core/config"
	"ephemeris/core/middleware"
	"ephemeris/core/routes"
	"net/url"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func Configure(ephemeris config.EphemerisConfig, m *martini.ClassicMartini) error {
	var store sessions.Store

	if ephemeris.Environment == "production" {
		m.Use(middleware.Syslog(middleware.SyslogOptions{
			Server: ephemeris.Syslog.Server,
		}))

		uri, err := url.Parse(ephemeris.Session.Redis.URL)

		if err != nil {
			return err
		}

		password := ""

		if uri.User != nil {
			password, _ = uri.User.Password()
		}

		store, err = sessions.NewRediStore(
			ephemeris.Session.Redis.MaxIdleConnections,
			uri.Scheme,
			uri.Host,
			password,
			ephemeris.Session.KeyPairs...,
		)

		if err != nil {
			return err
		}
	} else {
		m.Use(middleware.Stdout())
		store = sessions.NewCookieStore(ephemeris.Session.KeyPairs...)
	}

	m.Use(sessions.Sessions(ephemeris.Session.Name, store))
	m.Use(render.Renderer())
	m.Use(middleware.Database(middleware.DBOptions{
		Driver:             ephemeris.Database.Driver,
		URL:                ephemeris.Database.URL,
		MaxIdleConnections: ephemeris.Database.MaxIdleConnections,
	}))

	if os.Getenv("DEV_RUNNER") == "1" {
		m.Use(middleware.Fresh)
	}

	m.Group(ephemeris.APIRoot, func(r martini.Router) {
		routes.Apply(r)
	})

	return nil
}
