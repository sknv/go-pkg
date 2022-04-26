package render

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	type (
		input struct {
			status int
			data   any
		}

		result struct {
			code     int
			response string
		}
	)

	tests := map[string]struct {
		input input
		want  result
	}{
		"renders the provided code": {
			input: input{
				status: http.StatusCreated,
				data:   nil,
			},
			want: result{
				code:     http.StatusCreated,
				response: "null",
			},
		},
		"renders the provided data": {
			input: input{
				status: http.StatusOK,
				data:   map[string]any{"key": "val"},
			},
			want: result{
				code:     http.StatusOK,
				response: `{"key":"val"}`,
			},
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := httptest.NewRecorder()

			JSON(res, tc.input.status, tc.input.data)
			assert.Equal(t, tc.want.code, res.Code)
			assert.Equal(t, tc.want.response, strings.TrimSpace(res.Body.String()))
		})
	}
}
