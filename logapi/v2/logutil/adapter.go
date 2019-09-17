package logutil

import (
	"github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v2"
)

// NewLogCreating create a line number printing version of logapi.LogCreating
func NewLogCreating(logLevel logapi.Level, writer LogWriting) logapi.LogCreatingV2 {
	return logCreater{
		logLevel: logLevel,
		writer:   writer,
	}
}

type logCreater struct {
	logLevel logapi.Level
	writer   LogWriting
}

// Info prints an info msg to log.
func (logr logCreater) Info(msg logapi.LogAwareStringing, parms ...interface{}) {
	if logr.logLevel > logapi.InfoLogging {
		return
	}
	logr.writer.Write(msg.LogAwareToString(logapi.InfoLogging, parms))
}

// Error prints an error msg to log.
func (logr logCreater) Error(msg logapi.LogAwareStringing, parms ...interface{}) {
	if logr.logLevel > logapi.ErrorLogging {
		return
	}
	logr.writer.Write(msg.LogAwareToString(logapi.ErrorLogging, parms))
}

// Warn prints an warn msg to log.
func (logr logCreater) Warn(msg logapi.LogAwareStringing, parms ...interface{}) {
	if logr.logLevel > logapi.WarnLogging {
		return
	}
	logr.writer.Write(msg.LogAwareToString(logapi.WarnLogging, parms))
}

// Debug prints an debug msg to log.
func (logr logCreater) Debug(msg logapi.LogAwareStringing, parms ...interface{}) {
	if logr.logLevel > logapi.DebugLogging {
		return
	}
	logr.writer.Write(msg.LogAwareToString(logapi.DebugLogging, parms))
}
