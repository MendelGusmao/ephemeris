package middleware

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Authorize() martini.Handler {
	return func(logger *ApplicationLogger, renderer render.Render, session sessions.Session) {
		if session.Get("user.id") == nil {
			logger.Log("Not allowed")
			renderer.Status(http.StatusForbidden)
			return
		}
	}
}

func AuthorizeAdministrator() martini.Handler {
	return func(logger *ApplicationLogger, renderer render.Render, session sessions.Session) {
		administrator := session.Get("user.administrator")

		if session.Get("user.id") == nil || (administrator != nil && !administrator.(bool)) {
			logger.Log("Not allowed")
			renderer.Status(http.StatusForbidden)
			return
		}
	}
}
