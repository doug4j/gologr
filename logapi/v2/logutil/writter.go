package logutil

type LogWriting interface {
	Write(msg string)
}
