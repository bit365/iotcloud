package logger

import (
	"log"
	"os"
)

// LogLevel defines the level of logging
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

var (
	logger   *log.Logger
	logLevel LogLevel
)

// Init initializes the logger with a specific log level
func Init(level LogLevel) {
	logLevel = level
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an informational message
func Info(format string, v ...interface{}) {
	if logLevel <= INFO {
		logger.SetPrefix("INFO: ")
		logger.Printf(format, v...)
	}
}

// Warn logs a warning message
func Warn(format string, v ...interface{}) {
	if logLevel <= WARN {
		logger.SetPrefix("WARN: ")
		logger.Printf(format, v...)
	}
}

// Error logs an error message
func Error(format string, v ...interface{}) {
	if logLevel <= ERROR {
		logger.SetPrefix("ERROR: ")
		logger.Printf(format, v...)
	}
}
