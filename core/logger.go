package core

import "log/syslog"

type LogPriority syslog.Priority

type Logger interface {
	Log(priority syslog.Priority, message interface{}) error
	Logf(priority syslog.Priority, format string, message ...interface{}) error
}

func (p LogPriority) String() string {
	switch syslog.Priority(p) {
	case syslog.LOG_EMERG:
		return "EMERG"
	case syslog.LOG_ALERT:
		return "ALERT"
	case syslog.LOG_CRIT:
		return "CRIT"
	case syslog.LOG_ERR:
		return "ERR"
	case syslog.LOG_WARNING:
		return "WARNING"
	case syslog.LOG_NOTICE:
		return "NOTICE"
	case syslog.LOG_INFO:
		return "INFO"
	case syslog.LOG_DEBUG:
		return "DEBUG"
	}

	return ""
}
