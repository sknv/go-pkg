package problem

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	type result struct {
		code     int
		response string
	}

	tests := map[string]struct {
		err  error
		want result
	}{
		"renders a problem with proper status and default type": {
			err: New(http.StatusCreated, "any"),
			want: result{
				code:     http.StatusCreated,
				response: `{"type":"about:blank","title":"any","status":201}`,
			},
		},
		"renders a problem with proper status and type": {
			err: New(http.StatusOK, "any-title").SetType("any-type"),
			want: result{
				code:     http.StatusOK,
				response: `{"type":"any-type","title":"any-title","status":200}`,
			},
		},
		"renders any error as a problem with the default status": {
			err: errors.New("any"),
			want: result{
				code:     _defaultHTTPStatus,
				response: `{"type":"about:blank","title":"any","status":500}`,
			},
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := httptest.NewRecorder()

			Render(res, tc.err)
			assert.Equal(t, tc.want.code, res.Code)
			assert.Equal(t, tc.want.response, strings.TrimSpace(res.Body.String()))
		})
	}
}
