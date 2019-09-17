package logapi

// Logging generates log messages
type Logging interface {
	Info(msg string)
	Error(msg string)
	Warn(msg string)
	Debug(msg string)
	LogLevel() Level
}
