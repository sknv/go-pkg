package log

import (
	"context"
)

type ctxMarkerLogger struct{}

type ctxLogger struct {
	logger Logger
	fields Fields
}

var (
	_ctxKeyLogger = ctxMarkerLogger{}
)

// AddFields adds logger fields to the context.
func AddFields(ctx context.Context, fields Fields) {
	log, ok := ctx.Value(_ctxKeyLogger).(*ctxLogger)
	if !ok || log == nil {
		return
	}

	for k, v := range fields {
		log.fields[k] = v
	}
}

// Extract returns the logger with provided fields.
func Extract(ctx context.Context) Logger {
	log, ok := ctx.Value(_ctxKeyLogger).(*ctxLogger)
	if !ok || log == nil {
		return NullLogger
	}

	// Add log fields added until now
	fields := make(Fields, len(log.fields))
	for k, v := range log.fields {
		fields[k] = v
	}
	return log.logger.WithFields(fields)
}

// ToContext adds the logger to the context for extraction later, returning the new context that has been created.
func ToContext(ctx context.Context, logger Logger) context.Context {
	log := ctxLogger{
		logger: logger,
		fields: Fields{},
	}
	return context.WithValue(ctx, _ctxKeyLogger, &log)
}
