package logger

import (
	"log"

	"github.com/op/go-logging"
)

// Logger wraps a third-party logging struct.
type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
}

type gologger struct {
	*logging.Logger
}

type stdLogger struct{}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func (l *stdLogger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func (l *stdLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Error logs a message using ERROR as log level.
func (l *stdLogger) Error(args ...interface{}) {
	log.Print(args...)
}

// Errorf logs a message using ERROR as log level.
func (l *stdLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Warning logs a message using WARNING as log level.
func (l *stdLogger) Warning(args ...interface{}) {
	log.Print(args...)
}

// Warningf logs a message using WARNING as log level.
func (l *stdLogger) Warningf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Info logs a message using INFO as log level.
func (l *stdLogger) Info(args ...interface{}) {
	log.Print(args...)
}

// Infof logs a message using INFO as log level.
func (l *stdLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Debug logs a message using DEBUG as log level.
func (l *stdLogger) Debug(args ...interface{}) {
	log.Print(args...)
}

// Debugf logs a message using DEBUG as log level.
func (l *stdLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// New creates a new logger
func New(prefix string) Logger {
	l, err := logging.GetLogger(prefix)
	if err != nil {
		// use std logger
		return &stdLogger{}
	}

	return &gologger{
		Logger: l,
	}
}

func init() {
	formatter := logging.MustStringFormatter(`%{color}[%{module}] %{shortfunc} (%{shortfile}) [%{level}] ▶ %{color:reset} %{message}`)
	logging.SetFormatter(formatter)
}
