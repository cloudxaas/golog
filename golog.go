package cxlog

import "log"

const (
	Trace Level = iota
	Debug
	Info
	Notice
	Warn
	Error
	Fatal
)

func Display(logLevel, desiredLevel int, format string, args ...interface{}) {
	if desiredLevel >= logLevel {
		log.Printf(format, args...)
	}
}
