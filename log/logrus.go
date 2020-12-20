package log

import (
	"log"

	"github.com/sirupsen/logrus"
)

type Logger logrus.FieldLogger

const (
	DefaultLevel = logrus.InfoLevel
)

func ParseLevel(level string) logrus.Level {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = DefaultLevel
	}
	return lvl
}

const (
	JSONFormatter = "json"
	TextFormatter = "text"
)

var formatters = map[string]logrus.Formatter{
	JSONFormatter: &logrus.JSONFormatter{},
	TextFormatter: &logrus.TextFormatter{},
}

func ParseFormatter(formatter string) logrus.Formatter {
	if fmt, ok := formatters[formatter]; ok {
		return fmt
	}
	return formatters[JSONFormatter] // default formatter
}

// Option configures *logrus.Logger.
type Option func(*logrus.Logger)

// Build a log instance.
func Build(level, formatter string, options ...Option) Logger {
	logger := logrus.New()
	logger.SetLevel(ParseLevel(level))
	logger.SetFormatter(ParseFormatter(formatter))
	log.SetOutput(logger.Writer()) // redirect std log output

	// Apply options
	for _, opt := range options {
		opt(logger)
	}
	return logger
}
