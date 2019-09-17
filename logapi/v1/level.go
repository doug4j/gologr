package logapi

import (
	"errors"
	"strings"
)

// TODO(doug4j@gmail.com): More strongly type the logging options

// Level is the state in which logging is valid
type Level int

const (
	// DebugLogging is the lowest logging option
	DebugLogging Level = iota
	// InfoLogging is the logging option above info
	InfoLogging
	// WarnLogging is the logging option above info
	WarnLogging
	// ErrorLogging is the logging option above warn
	ErrorLogging
	// FatalLogging is the logging option above error
	FatalLogging
)

// IntP returns the uint for level
func (l Level) IntP() *uint {
	myInt := uint(l)
	return &myInt
}

func (l Level) String() string {
	switch l {
	case InfoLogging:
		return "INFO"
	case WarnLogging:
		return "WARN"
	case ErrorLogging:
		return "ERROR"
	case FatalLogging:
		return "FATAL"
	default:
		//	case DebugLogging:
		return "DEBUG"
	}
}

func ParseLevel(val string) (Level, error) {
	val = strings.ToLower(val)
	val = strings.TrimSpace(val)
	switch val {
	case "info":
		return InfoLogging, nil
	case "warn":
		return WarnLogging, nil
	case "error":
		return ErrorLogging, nil
	case "fatal":
		return FatalLogging, nil
	case "debug":
		return DebugLogging, nil
	default:
		return DebugLogging, errors.New("cannot parse value [" + val + "]")
	}
}
