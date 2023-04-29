package cxlog

import "log"

const (
	Trace = iota
	Debug
	Info
	Notice
	Warn
	Error
	Fatal
)

func Display(logLevel, desiredLevel uint8, format string, args ...interface{}) {
	if desiredLevel >= logLevel {
		log.Printf(format, args...)
	}
}
