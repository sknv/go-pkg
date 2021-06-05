package log

import (
	"io/ioutil"
	"log"

	"github.com/sirupsen/logrus"
)

// Proxy
type (
	FieldLogger = logrus.FieldLogger
	Logger      = logrus.Logger
	Fields      = logrus.Fields
)

const DefaultLevel = logrus.InfoLevel

func ParseLevel(level string) logrus.Level {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = DefaultLevel
	}
	return lvl
}

type Formatter string

const (
	JSONFormatter Formatter = "json"
	TextFormatter Formatter = "text"
)

var formatters = map[Formatter]logrus.Formatter{
	JSONFormatter: &logrus.JSONFormatter{},
	TextFormatter: &logrus.TextFormatter{},
}

func ParseFormatter(formatter Formatter) logrus.Formatter {
	if fmt, ok := formatters[formatter]; ok {
		return fmt
	}
	return formatters[JSONFormatter] // default formatter
}

var NullLogger = &logrus.Logger{
	Out:       ioutil.Discard,
	Formatter: &logrus.TextFormatter{},
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.PanicLevel,
}

type Config struct {
	Formatter Formatter `mapstructure:"formatter"`
	Level     string    `mapstructure:"level"`
}

// Option configures *Logger.
type Option func(*Logger)

// Build a logger instance.
func Build(config Config, options ...Option) *Logger {
	logger := logrus.New()
	logger.SetLevel(ParseLevel(config.Level))
	logger.SetFormatter(ParseFormatter(config.Formatter))
	log.SetOutput(logger.Writer()) // redirect std log output

	// Apply options
	for _, opt := range options {
		opt(logger)
	}
	return logger
}
