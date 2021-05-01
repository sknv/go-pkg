package postgres

import (
	"context"
	stdlog "log"

	"github.com/jackc/pgx/v4"
	"github.com/sknv/go-pkg/log"
)

func WithLogger(level string) Option {
	logLevel, err := pgx.LogLevelFromString(level)
	if err != nil {
		stdlog.Printf("pgx.LogLevelFromString: %s", err)
		logLevel = pgx.LogLevelNone // fallback level
	}

	return func(cfg *pgx.ConnConfig) {
		cfg.Logger = &Logger{}
		cfg.LogLevel = logLevel
	}
}

type Logger struct{}

func (l *Logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	logger := log.Extract(ctx)
	if data != nil {
		logger = logger.WithFields(data)
	}

	switch level {
	case pgx.LogLevelTrace:
		logger.WithField("PGX_LOG_LEVEL", level).Debug(msg)
	case pgx.LogLevelDebug:
		logger.Debug(msg)
	case pgx.LogLevelInfo:
		logger.Info(msg)
	case pgx.LogLevelWarn:
		logger.Warn(msg)
	case pgx.LogLevelError:
		logger.Error(msg)
	default:
		logger.WithField("INVALID_PGX_LOG_LEVEL", level).Error(msg)
	}
}
