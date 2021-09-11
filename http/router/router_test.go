package router

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		opt            func(*bool) Option
		optApplied     bool
		wantOptApplied bool
	}{
		"applies options successfully": {
			opt: func(opt *bool) Option {
				return func(*chi.Mux) { *opt = true }
			},
			wantOptApplied: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_ = New(tc.opt(&tc.optApplied))
			assert.Equal(t, tc.wantOptApplied, tc.optApplied)
		})
	}
}
