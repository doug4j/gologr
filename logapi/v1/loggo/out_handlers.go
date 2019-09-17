package loggo

import (
	"fmt"
	"log"
)

func NewLogOutPrintln() OutputHandler {
	var out = func(msg string) {
		log.Println(msg)
	}
	return out
}

func NewStdOutPrintln() OutputHandler {
	var out = func(msg string) {
		fmt.Println(msg)
	}
	return out
}
