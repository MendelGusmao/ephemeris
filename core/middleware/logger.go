package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-martini/martini"
)

var (
	layout = "02/01/2006 15:04:05"
	stdout = log.New(os.Stdout, "[ephemeris] ", 0)
)

type AppLogger struct {
	*log.Logger
	Request *http.Request
}

func (logger *AppLogger) Log(message string) {
	addr := logger.Request.Header.Get("X-Real-IP")
	if addr == "" {
		addr = logger.Request.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = logger.Request.RemoteAddr
		}
	}

	now := time.Now().Format(layout)
	logger.Printf("%s %s for %s @ %s -> %s", logger.Request.Method, logger.Request.URL.Path, addr, now, message)
}

func (logger *AppLogger) Logf(format string, parts ...interface{}) {
	logger.Log(fmt.Sprintf(format, parts...))
}

func Logger() martini.Handler {
	return func(c martini.Context, req *http.Request) {
		c.Map(&AppLogger{
			Request: req,
			Logger:  stdout,
		})

		c.Next()
	}
}
