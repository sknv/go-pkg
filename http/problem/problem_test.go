package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetType(t *testing.T) {
	tests := map[string]struct {
		problemType string
		want        string
	}{
		"returns the default type if one is not set": {problemType: "", want: _defaultType},
		"returns the provided type":                  {problemType: "any", want: "any"},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			prb := New(0, "").SetType(tc.problemType)
			assert.Equal(t, tc.want, prb.GetType())
		})
	}
}

func TestError(t *testing.T) {
	tests := map[string]struct {
		status int
		title  string
		detail string
		want   string
	}{
		"returns the expected error": {
			status: 200,
			title:  "any-title",
			detail: "any-detail",
			want:   "status = 200, title = any-title, detail = any-detail",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			prb := New(tc.status, tc.title).SetDetail(tc.detail)
			assert.Equal(t, tc.want, prb.Error())
		})
	}
}
