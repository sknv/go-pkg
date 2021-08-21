package closer

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	err1, err2 := errors.New("1"), errors.New("2")

	tests := map[string]struct {
		closers []Closer
		want    error
	}{
		"closes successfully": {
			closers: []Closer{
				func(context.Context) error {
					return nil
				},
			},
		},
		"closes with a single error": {
			closers: []Closer{
				func(context.Context) error {
					return nil
				},
				func(context.Context) error {
					return err1
				},
			},
			want: multierror.Append(err1, nil),
		},
		"closes with multiple errors in correct order": {
			closers: []Closer{
				func(context.Context) error {
					return err1
				},
				func(context.Context) error {
					return err2
				},
			},
			want: multierror.Append(err2, err1),
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			closers := NewClosers()
			for _, cls := range tc.closers {
				closers.Add(cls)
			}

			err := closers.Close(context.Background())
			assert.Equal(t, tc.want, err)
		})
	}
}
