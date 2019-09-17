package loggo

import (
	logapi "github.com/doug4j/gologr/logapi/v1"
)

func NewDefaultLogging(logLevel logapi.Level) logapi.Logging {
	if logLevel == logapi.InfoLogging {
		return NewLogAdaptor(logLevel,
			NewEmojiMessageHandler(),
			NewStdOutPrintln())
	}
	return newDebugLogging(logLevel)
}

func newDebugLogging(logLevel logapi.Level) logapi.Logging {
	return NewLogAdaptor(logLevel,
		NewLineNumMessageHandler(LineNumMessageHandlerParms{PadLogLevel: true}),
		NewStdOutPrintln())
}

// // NewLogCreating create a line number printing version of logapi.LogCreating
// func NewLogCreating(logLevel logapi.Level) logapi.LogCreating {
// 	return logCreater{
// 		logLevel: logLevel,
// 	}
// }

// type logCreater struct {
// 	logLevel logapi.Level
// }

// // LogLevel returns the configured currrent log level for the underlying logger.
// func (logr logCreater) LogLevel() logapi.Level {
// 	return logr.logLevel
// }

// // Info prints an info msg to log.
// func (logr logCreater) Info(msg string) {
// 	if logr.logLevel > logapi.InfoLogging {
// 		return
// 	}
// 	pc, file, line, _ := runtime.Caller(1)

// 	logWithLastFunc(pc, line, file, msg, "INFO")
// }

// // Error prints an error msg to log.
// func (logr logCreater) Error(msg string) {
// 	if logr.logLevel > logapi.ErrorLogging {
// 		return
// 	}
// 	pc, file, line, _ := runtime.Caller(1)

// 	logWithLastFunc(pc, line, file, msg, "ERROR")
// }

// // Warn prints an warn msg to log.
// func (logr logCreater) Warn(msg string) {
// 	if logr.logLevel > logapi.WarnLogging {
// 		return
// 	}
// 	pc, file, line, _ := runtime.Caller(1)

// 	logWithLastFunc(pc, line, file, msg, "WARN")
// }

// // Debug prints an debug msg to log.
// func (logr logCreater) Debug(msg string) {
// 	if logr.logLevel > logapi.DebugLogging {
// 		return
// 	}
// 	pc, file, line, _ := runtime.Caller(1)
// 	logWithLastFunc(pc, line, file, msg, "DEBUG")
// }

// func logWithLastFunc(pc uintptr, line int, file, msg, prefix string) (justPcName, justFileName string) {
// 	fullPCName := runtime.FuncForPC(pc).Name()
// 	// lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
// 	// justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]
// 	justPcName = fullPCName

// 	// lastIndexOfFile := strings.LastIndex(file, "/") + 1
// 	// justFileName := file[lastIndexOfFile:len(file)]
// 	justFileName = file

// 	// log.Printf("%s [%s:%d] [%s] %v", prefix, justFileName, line, justPcName, msg)
// 	log.Printf("%s [%s:%d]\n%v", prefix, justFileName, line, msg)

// 	return justPcName, justFileName
// }
