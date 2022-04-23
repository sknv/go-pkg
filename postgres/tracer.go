package postgres

import (
	"github.com/XSAM/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// RegisterTracer registers tracing for the provided driver name.
func RegisterTracer(driverName string) (string, error) {
	return otelsql.Register(driverName, otelsql.WithAttributes(
		semconv.DBSystemPostgreSQL,
	))
}
