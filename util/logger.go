package util

import (
	clog "github.com/charmbracelet/log"
	"os"
	"strings"
)

// LoggerFunc represents a function that can be used for logging.
// It takes an interface{} argument which represents the log level or message,
// and a variadic ...interface{} argument for additional log data.
type LoggerFunc func(interface{}, ...interface{})

// loggerFunc is a variable of type LoggerFunc, which represents a function that can be used for logging.
// It takes an interface{} argument which represents the log level or message,
// and a variadic ...interface{} argument for additional log data.
var loggerFunc LoggerFunc = nil

// SetLoggerFunc sets the global logger function to the provided LoggerFunc.
// The global logger function is used for logging messages throughout the library.
func SetLoggerFunc(l LoggerFunc) {
	loggerFunc = l
}

// Debug logs a debug message using the global logger function if it is set.
// The debug message can be any type. Additional arguments can be provided.
// If loggerFunc is nil, no action is taken.
func Debug(message interface{}, args ...interface{}) {
	if loggerFunc != nil {
		loggerFunc(message, args...)
	}
}

// wantDefaultLoggerFunc checks whether the environment variable "OYSTER_DEBUG" is set to a non-zero value.
// Returns true if the variable is set to a non-zero value, false otherwise.
func wantDefaultLoggerFunc() bool {
	if val, ok := os.LookupEnv("OYSTER_DEBUG"); ok {
		val = strings.ToLower(strings.TrimSpace(val))
		return val != "0" && val != "no" && val != "off"
	} else {
		return false
	}
}

// setDefaultLoggerFunc initializes a default logger with specified options and sets it as the logger function.
func setDefaultLoggerFunc() {
	logger := clog.NewWithOptions(os.Stderr, clog.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Prefix:          "🦪 oyster",
	})
	logger.SetLevel(clog.DebugLevel)
	SetLoggerFunc(logger.Debug)
}

func init() {
	if wantDefaultLoggerFunc() {
		setDefaultLoggerFunc()
	}
}
