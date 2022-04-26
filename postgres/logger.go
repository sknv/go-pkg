package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/sknv/go-pkg/log"
)

const (
	DefaultLogLevel = pgx.LogLevelNone
)

func WithLogger(level string) Option {
	logLevel, err := pgx.LogLevelFromString(level)
	if err != nil {
		log.L().WithError(err).Warn("parse pgx log level")
		logLevel = DefaultLogLevel
	}

	return func(cfg *pgx.ConnConfig) {
		cfg.Logger = logger{}
		cfg.LogLevel = logLevel
	}
}

type logger struct{}

func (logger) Log(ctx context.Context, _ pgx.LogLevel, msg string, data map[string]any) {
	log.Extract(ctx).WithField("op", "postgres").WithFields(data).Debug(msg) // always use debug level
}
