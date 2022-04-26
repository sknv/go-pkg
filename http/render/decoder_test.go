package render

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeJSON(t *testing.T) {
	tests := map[string]struct {
		input   string
		wantErr bool
	}{
		"returns error if an input is invalid": {input: `{"key": "val}`, wantErr: true},
		"decodes successfully":                 {input: `{"key": "val"}`, wantErr: false},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dst := make(map[string]any)
			err := DecodeJSON(strings.NewReader(tc.input), &dst)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
