package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultConfig(t *testing.T) {
	tests := map[string]struct {
		wantSuccesses bool
	}{
		"creates default config with expected options successfully": {wantSuccesses: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg := NewDefaultConfig()
			assert.Equal(t, tc.wantSuccesses, cfg.Producer.Return.Successes)
		})
	}
}
