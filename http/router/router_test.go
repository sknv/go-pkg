package router

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var optApplied bool
	tests := map[string]struct {
		opt            Option
		wantOptApplied bool
	}{
		"applies options successfully": {
			opt:            func(*chi.Mux) { optApplied = true },
			wantOptApplied: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			optApplied = false // restore every time

			_ = New(tc.opt)
			assert.Equal(t, tc.wantOptApplied, optApplied)
		})
	}
}
