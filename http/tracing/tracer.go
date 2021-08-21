package tracing

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/sknv/go-pkg/log"
	"github.com/sknv/go-pkg/tracing"
)

const (
	_fieldTraceID = "trace_id"
)

// WithTracer middleware enables tracing and injects a trace id into the context of each request.
func WithTracer(operation string, options ...otelhttp.Option) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			traceID, ok := tracing.GetTraceID(r.Context())
			if ok {
				log.AddFields(r.Context(), logrus.Fields{_fieldTraceID: traceID})
			}
			next.ServeHTTP(w, r)
		}

		handler := http.HandlerFunc(fn)
		return otelhttp.NewHandler(handler, operation, options...)
	}
}

// ForHTTPClient wraps the provided http client with a tracing transport.
func ForHTTPClient(client *http.Client, options ...otelhttp.Option) *http.Client {
	newClient := *client                                                         // copy the existing client
	newClient.Transport = otelhttp.NewTransport(newClient.Transport, options...) // wrap transport
	return &newClient
}
