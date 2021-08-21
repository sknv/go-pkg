package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("inits logger with an empty config successfully", func(t *testing.T) {
		Init(Config{})
		assert.NotEqual(t, L(), _nullLogger)
	})
}

func TestParseLevel(t *testing.T) {
	tests := map[string]struct {
		level string
		want  logrus.Level
	}{
		"parses known level successfully":           {level: "debug", want: logrus.DebugLevel},
		"parses unknown level with the default one": {level: "unknown", want: DefaultLevel},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			lvl := ParseLevel(tc.level)
			assert.Equal(t, tc.want, lvl)
		})
	}
}

func TestBuild(t *testing.T) {
	var optApplied bool
	tests := map[string]struct {
		level          logrus.Level
		opts           []Option
		wantOptApplied bool
	}{
		"builds logger successfully": {
			level:          logrus.DebugLevel,
			wantOptApplied: false,
		},
		"applies options successfully": {
			level:          logrus.DebugLevel,
			opts:           []Option{func(*logrus.Logger) { optApplied = true }},
			wantOptApplied: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			optApplied = false // restore every time

			Build(tc.level, tc.opts...)
			assert.NotEqual(t, logrus.StandardLogger(), _nullLogger)
			assert.Equal(t, tc.wantOptApplied, optApplied)
		})
	}
}
