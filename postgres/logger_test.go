package postgres

import (
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestWithLogger(t *testing.T) {
	tests := map[string]struct {
		level string
		want  pgx.LogLevel
	}{
		"parses valid log level successfully": {
			level: "debug",
			want:  pgx.LogLevelDebug,
		},
		"sets fallback log level if one is not defined": {
			level: "invalid",
			want:  DefaultLogLevel,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var cfg pgx.ConnConfig
			WithLogger(tc.level)(&cfg)
			assert.Equal(t, tc.want, cfg.LogLevel)
		})
	}
}
