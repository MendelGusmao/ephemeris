package middleware

import (
	"events/lib/martini"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	layout = "02/01/2006 15:04:05"
)

type ApplicationLogger struct {
	*log.Logger
	*http.Request
}

func (logger *ApplicationLogger) Log(message string) {
	addr := logger.Request.Header.Get("X-Real-IP")
	if addr == "" {
		addr = logger.Request.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = logger.Request.RemoteAddr
		}
	}

	now := time.Now().Format(layout)
	logger.Logger.Printf("%s %s for %s @ %s -> %s", logger.Request.Method, logger.Request.URL.Path, addr, now, message)
}

func Logger() martini.Handler {
	return func(c martini.Context, req *http.Request) {
		c.Map(&ApplicationLogger{
			Request: req,
			Logger:  log.New(os.Stdout, "[events] ", 0),
		})

		c.Next()
	}
}
