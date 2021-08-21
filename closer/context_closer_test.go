package closer

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseContext(t *testing.T) {
	doneCtx, cancel := context.WithTimeout(context.Background(), -1)
	defer cancel()

	tests := map[string]struct {
		ctx     context.Context
		closer  PlainCloser
		wantErr bool
	}{
		"ignores empty closer": {
			ctx:     context.Background(),
			closer:  nil,
			wantErr: false,
		},
		"closes successfully": {
			ctx:     context.Background(),
			closer:  func() error { return nil },
			wantErr: false,
		},
		"closer returns an error": {
			ctx:     context.Background(),
			closer:  func() error { return errors.New("any") },
			wantErr: true,
		},
		"returns an error for done context": {
			ctx:     doneCtx,
			closer:  func() error { return nil },
			wantErr: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := CloseContext(tc.ctx, tc.closer)
			assert.Equal(t, tc.wantErr, err != nil, "errors do not match")
		})
	}
}
