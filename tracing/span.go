package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func GetTraceID(ctx context.Context) (string, bool) {
	span := trace.SpanFromContext(ctx)
	return span.SpanContext().TraceID().String(), span.SpanContext().TraceID().IsValid()
}
