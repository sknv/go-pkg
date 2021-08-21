package tracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestInit(t *testing.T) {
	t.Run("inits tracer with an empty config successfully", func(t *testing.T) {
		err := Init(Config{})
		assert.NoError(t, err)
	})
}

func TestNewJaegerTracerProvider(t *testing.T) {
	t.Run("creates tracer provider an empty config successfully", func(t *testing.T) {
		_, err := NewJaegerTracerProvider(Config{})
		assert.NoError(t, err)
	})
}

func TestInitTracer(t *testing.T) {
	dummyProvider, _ := NewJaegerTracerProvider(Config{})

	tests := map[string]struct {
		provider trace.TracerProvider
		wantErr  bool
	}{
		"returns an error for an empty tracer provider":              {provider: nil, wantErr: true},
		"inits tracer with a non-empty tracer provider successfully": {provider: dummyProvider, wantErr: false},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := InitTracer(tc.provider)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
