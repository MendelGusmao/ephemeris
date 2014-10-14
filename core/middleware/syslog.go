package middleware

import (
	"ephemeris/core"
	"ephemeris/core/config"
	"fmt"
	"log/syslog"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

var (
	_ core.Logger = (*SysLog)(nil)
)

type SysLog struct {
	*syslog.Writer
	template string
	level    syslog.Priority
}

type SyslogOptions struct {
	URL string
}

func Syslog(syslogOptions SyslogOptions, level syslog.Priority) martini.Handler {
	writer, err := syslog.Dial("tcp", syslogOptions.URL, syslog.LOG_INFO, "ephemeris")

	if err != nil {
		writer, err = syslog.Dial("udp", syslogOptions.URL, syslog.LOG_INFO, "ephemeris")

		if err != nil {
			return Stdout(level)
		}
	}

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

		template := fmt.Sprintf("%s %s for %s -> %%v",
			req.Method,
			req.URL.Path,
			addr,
		)

		sysLog := SysLog{
			Writer:   writer,
			template: template,
			level:    level,
		}

		sysLog.Log(syslog.LOG_DEBUG, "Started")

		c.Map(&sysLog)
		rw := res.(martini.ResponseWriter)
		c.Next()

		sysLog.Logf(syslog.LOG_DEBUG, "Completed %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), time.Since(start))
	}
}

func (logger *SysLog) Log(priority syslog.Priority, message interface{}) error {
	if priority > config.Ephemeris.Log.Level {
		return nil
	}

	msg := fmt.Sprintf(logger.template, message)

	switch priority {
	case syslog.LOG_EMERG:
		return logger.Writer.Emerg(msg)
	case syslog.LOG_ALERT:
		return logger.Writer.Alert(msg)
	case syslog.LOG_CRIT:
		return logger.Writer.Crit(msg)
	case syslog.LOG_ERR:
		return logger.Writer.Err(msg)
	case syslog.LOG_WARNING:
		return logger.Writer.Warning(msg)
	case syslog.LOG_NOTICE:
		return logger.Writer.Notice(msg)
	case syslog.LOG_INFO:
		return logger.Writer.Info(msg)
	case syslog.LOG_DEBUG:
		return logger.Writer.Debug(msg)
	}

	return nil
}

func (logger *SysLog) Logf(priority syslog.Priority, format string, message ...interface{}) error {
	if priority > config.Ephemeris.Log.Level {
		return nil
	}

	return logger.Log(priority, fmt.Sprintf(format, message...))
}
