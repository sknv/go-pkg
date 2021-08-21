package tracing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func TestWithTracer(t *testing.T) {
	tests := map[string]struct {
		wantNil bool
	}{
		"wraps with tracer successfully": {wantNil: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h := WithTracer("any")
			assert.Equal(t, tc.wantNil, h == nil, "result does not match the expected one")
			h(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		})
	}
}

func TestForHttpClient(t *testing.T) {
	tests := map[string]struct {
		client *http.Client
	}{
		"wraps client's transport": {client: &http.Client{}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := ForHTTPClient(tc.client)
			_, ok := client.Transport.(*otelhttp.Transport)
			assert.True(t, ok, "new client's transport must be wrapped")

			_, ok = tc.client.Transport.(*otelhttp.Transport)
			assert.False(t, ok, "old client's transport must not be wrapped")
		})
	}
}
