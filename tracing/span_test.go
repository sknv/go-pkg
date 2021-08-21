package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTraceID(t *testing.T) {
	tests := map[string]struct {
		input  context.Context
		wantOK bool
	}{
		"returns an invalid flag for an empty span": {
			input:  context.Background(),
			wantOK: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, ok := GetTraceID(tc.input)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}
