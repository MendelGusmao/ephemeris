package middleware

import (
	"ephemeris/core"
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"time"

	"github.com/go-martini/martini"
)

var (
	layout             = "02/01/2006 15:04:05"
	stdout             = log.New(os.Stdout, "[ephemeris] ", 0)
	_      core.Logger = (*StdOut)(nil)
)

type StdOut struct {
	*log.Logger
	template string
}

func Stdout() martini.Handler {
	return func(c martini.Context, req *http.Request) {
		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
				if addr == "" {
					addr = "none"
				}
			}
		}

		template := fmt.Sprintf("[%%s] %s %s for %s @ %%s -> %%v",
			req.Method,
			req.URL.Path,
			addr,
		)

		c.Map(&StdOut{
			Logger:   stdout,
			template: template,
		})

		c.Next()
	}
}

func (logger *StdOut) Log(priority syslog.Priority, message interface{}) error {
	now := time.Now().Format(layout)
	logger.Printf(logger.template, core.LogPriority(priority), now, message)
	return nil
}

func (logger *StdOut) Logf(priority syslog.Priority, format string, message interface{}) error {
	return logger.Log(priority, fmt.Sprintf(format, message))
}
