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
	level    syslog.Priority
}

func Stdout(level syslog.Priority) martini.Handler {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request) {
		start := time.Now()

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

		stdout := StdOut{
			Logger:   stdout,
			template: template,
			level:    level,
		}

		stdout.Log(syslog.LOG_DEBUG, "Started")

		c.Map(&stdout)
		rw := res.(martini.ResponseWriter)
		c.Next()

		stdout.Logf(syslog.LOG_DEBUG, "Completed %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), time.Since(start))
	}
}

func (logger *StdOut) Log(priority syslog.Priority, message interface{}) error {
	if priority > logger.level {
		return nil
	}

	now := time.Now().Format(layout)
	logger.Printf(logger.template, core.LogPriority(priority), now, message)

	return nil
}

func (logger *StdOut) Logf(priority syslog.Priority, format string, message ...interface{}) error {
	if priority > logger.level {
		return nil
	}

	return logger.Log(priority, fmt.Sprintf(format, message...))
}
