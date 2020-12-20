package log

import (
	"context"
)

type ctxMarkerRequestID struct{}

var (
	_ctxKeyRequestID = ctxMarkerRequestID{}
)

// Extract returns the logger with a request id if exists.
func Extract(ctx context.Context, logger Logger) Logger {
	if ctx == nil { // return the default logger if a context is nil
		return logger
	}

	requestID, _ := GetRequestID(ctx)
	if requestID == "" {
		return logger
	}
	return logger.WithField("requestID", requestID)
}

// GetRequestID retrieves request id from the context.
func GetRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(_ctxKeyRequestID).(string)
	return requestID, ok
}

// PutRequestID puts request id into the context.
func PutRequestID(ctx context.Context, requestID string) context.Context {
	if requestID == "" {
		return ctx
	}
	return context.WithValue(ctx, _ctxKeyRequestID, requestID)
}
