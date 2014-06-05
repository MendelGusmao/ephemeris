package middleware

import (
	"fmt"
	"github.com/go-martini/martini"
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
	request *http.Request
}

func (logger *ApplicationLogger) Log(message string) {
	addr := logger.request.Header.Get("X-Real-IP")
	if addr == "" {
		addr = logger.request.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = logger.request.RemoteAddr
		}
	}

	now := time.Now().Format(layout)
	logger.Logger.Printf("%s %s for %s @ %s -> %s", logger.request.Method, logger.request.URL.Path, addr, now, message)
}

func (logger *ApplicationLogger) Logf(format string, parts ...interface{}) {
	logger.Log(fmt.Sprintf(format, parts...))
}

func Logger() martini.Handler {
	return func(c martini.Context, req *http.Request) {
		c.Map(&ApplicationLogger{
			request: req,
			Logger:  log.New(os.Stdout, "[ephemeris] ", 0),
		})

		c.Next()
	}
}
