package loggo

import (
	"github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v1"
)

type MessageHandler func(msg string, logLevel logapi.Level) string
type OutputHandler func(msg string)

// NewLogAdaptor create logapi.LogCreating with configurable handlers
func NewLogAdaptor(logLevel logapi.Level, msgHandler MessageHandler, outHandler OutputHandler) logapi.Logging {
	return logAdaptor{
		logLevel:   logLevel,
		msgHandler: msgHandler,
		outHandler: outHandler,
	}
}

type logAdaptor struct {
	logLevel   logapi.Level
	msgHandler MessageHandler
	outHandler OutputHandler
}

// LogLevel returns the configured currrent log level for the underlying logger.
func (logr logAdaptor) LogLevel() logapi.Level {
	return logr.logLevel
}

// Info prints an info msg to log.
func (logr logAdaptor) Info(msg string) {
	if logr.logLevel > logapi.InfoLogging {
		return
	}
	logr.outHandler(logr.msgHandler(msg, logapi.InfoLogging))
}

// Error prints an error msg to log.
func (logr logAdaptor) Error(msg string) {
	if logr.logLevel > logapi.InfoLogging {
		return
	}
	logr.outHandler(logr.msgHandler(msg, logapi.ErrorLogging))
}

// Warn prints an warn msg to log.
func (logr logAdaptor) Warn(msg string) {
	if logr.logLevel > logapi.WarnLogging {
		return
	}
	logr.outHandler(logr.msgHandler(msg, logapi.WarnLogging))
}

// Debug prints an debug msg to log.
func (logr logAdaptor) Debug(msg string) {
	if logr.logLevel > logapi.DebugLogging {
		return
	}
	logr.outHandler(logr.msgHandler(msg, logapi.DebugLogging))
}
