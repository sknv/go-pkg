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
	type input struct {
		level logrus.Level
		opts  []func(*bool) Option
	}

	tests := map[string]struct {
		input          input
		optApplied     bool
		wantOptApplied bool
	}{
		"builds logger successfully": {
			input:          input{level: logrus.DebugLevel},
			wantOptApplied: false,
		},
		"applies options successfully": {
			input: input{
				level: logrus.DebugLevel,
				opts: []func(*bool) Option{
					func(opt *bool) Option {
						return func(*logrus.Logger) { *opt = true }
					},
				},
			},
			wantOptApplied: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			opts := make([]Option, 0, len(tc.input.opts))
			for _, opt := range tc.input.opts {
				opts = append(opts, opt(&tc.optApplied))
			}
			Build(tc.input.level, opts...)
			assert.NotEqual(t, logrus.StandardLogger(), _nullLogger)
			assert.Equal(t, tc.wantOptApplied, tc.optApplied)
		})
	}
}
