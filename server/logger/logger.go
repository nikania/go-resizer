package logger

import (
	"log"
	"os"
)

type Logger interface {
	Debug(v ...any)
	Info(v ...any)
	Warn(v ...any)
	Error(v ...any)
}

type FileLogger struct {
	logDebug, logInfo bool
	debug             *log.Logger
	info              *log.Logger
	warn              *log.Logger
	error             *log.Logger
}

type ConsoleLogger struct {
	logDebug, logInfo bool
	debug             *log.Logger
	info              *log.Logger
	warn              *log.Logger
	error             *log.Logger
}

func NewConsoleLogger(logDebug, logInfo bool) *ConsoleLogger {
	debug := log.New(os.Stdout, "DEBUG: ", log.Default().Flags())
	info := log.New(os.Stdout, "INFO: ", log.Default().Flags())
	warn := log.New(os.Stdout, "WARNING: ", log.Default().Flags())
	error := log.New(os.Stdout, "ERROR: ", log.Default().Flags())
	return &ConsoleLogger{
		logDebug, logInfo, debug, info, warn, error,
	}
}

func (l ConsoleLogger) Debug(v ...any) {
	if l.logDebug {
		l.debug.Println(v...)
	}
}

func (l ConsoleLogger) Info(v ...any) {
	if l.logInfo {
		l.info.Println(v...)
	}
}

func (l ConsoleLogger) Warn(v ...any) {
	l.warn.Println(v...)
}

func (l ConsoleLogger) Error(v ...any) {
	l.error.Println(v...)
}

func NewFileLogger(filename string, logDebug, logInfo bool) *FileLogger {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	debug := log.New(file, "DEBUG: ", log.Default().Flags())
	info := log.New(file, "INFO: ", log.Default().Flags())
	warn := log.New(file, "WARNING: ", log.Default().Flags())
	error := log.New(file, "ERROR: ", log.Default().Flags())
	return &FileLogger{
		logDebug, logInfo, debug, info, warn, error,
	}
}

func (l FileLogger) Debug(v ...any) {
	if l.logDebug {
		l.debug.Println(v...)
	}
}

func (l FileLogger) Info(v ...any) {
	if l.logInfo {
		l.info.Println(v...)
	}
}

func (l FileLogger) Warn(v ...any) {
	l.warn.Println(v...)
}

func (l FileLogger) Error(v ...any) {
	l.error.Println(v...)
}
