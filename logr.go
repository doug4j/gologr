package logr

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

//GetCallingName obtains the name of the calling function
func GetCallingName() string {
	pc, _, _, _ := runtime.Caller(1)

	fullPCName := runtime.FuncForPC(pc).Name()
	lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
	justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]
	lastIndexOfJustName := strings.LastIndex(justPcName, ".") + 1
	justName := justPcName[lastIndexOfJustName:len(justPcName)]

	return justName
}

//TODO: More strongly type the logging options

//LogLevel establishes the logging mode using the Level value.
var LogLevel = FatalLoging

//Level is the state in which logging is valid
type Level int

const (
	//DebugLogging is the lowest logging option
	DebugLogging Level = iota
	//InfoLogging is the logging option above info
	InfoLogging
	//WarnLogging is the logging option above info
	WarnLogging
	//ErrorLogging is the logging option above warn
	ErrorLogging
	//FatalLoging is the logging option above error
	FatalLoging
)

func (l Level) IntP() *uint {
	myInt := uint(l)
	return &myInt
}

func prepMsg(format interface{}, parms []interface{}) string {
	strFormt := fmt.Sprintf("%v", format)
	if len(parms) >= 1 {
		return fmt.Sprintf(strFormt, parms...)
	}
	return strFormt
}

func logWithLastFuncInfo(pc uintptr, file string, line int, format interface{}, parms []interface{}, prefix string) (string, string) {
	fullPCName := runtime.FuncForPC(pc).Name()
	lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
	justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]

	lastIndexOfFile := strings.LastIndex(file, "/") + 1
	justFileName := file[lastIndexOfFile:len(file)]

	log.Printf("%s [%s:%d] [%s] %v", prefix, justFileName, line, justPcName, prepMsg(format, parms))

	return justPcName, justFileName
}

//Info prints an info msg to log.
func Info(format interface{}, parms ...interface{}) {
	if LogLevel > InfoLogging {
		return
	}
	pc, file, line, _ := runtime.Caller(1)

	logWithLastFuncInfo(pc, file, line, format, parms, "INFO")
}

//Error prints an error msg to log.
func Error(format interface{}, parms ...interface{}) {
	if LogLevel > ErrorLogging {
		return
	}
	pc, file, line, _ := runtime.Caller(1)

	logWithLastFuncInfo(pc, file, line, format, parms, "ERROR")
}

//Warn prints an warn msg to log.
func Warn(format interface{}, parms ...interface{}) {
	if LogLevel > WarnLogging {
		return
	}
	pc, file, line, _ := runtime.Caller(1)

	logWithLastFuncInfo(pc, file, line, format, parms, "WARN")
}

//Debug prints an debug msg to log.
func Debug(format interface{}, parms ...interface{}) {
	if LogLevel > DebugLogging {
		return
	}
	pc, file, line, _ := runtime.Caller(1)
	logWithLastFuncInfo(pc, file, line, format, parms, "DEBUG")
}

//Fatal prints a fatal msg to log then panics.
func Fatal(format interface{}, parms ...interface{}) {
	//Always print fatal messages
	pc, file, line, _ := runtime.Caller(1)
	justPcName, justFileName := logWithLastFuncInfo(pc, file, line, format, parms, "FATAL")
	panic("Fatal Error: " + fmt.Sprintf("FATAL [%s:%d] [%s] %v", justFileName, line, justPcName, prepMsg(format, parms)))
}
