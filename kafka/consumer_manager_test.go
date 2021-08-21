package kafka

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumerManagerClose(t *testing.T) {
	doneCtx, cancel := context.WithTimeout(context.Background(), -1)
	defer cancel()

	tests := map[string]struct {
		ctx     context.Context
		wantErr bool
	}{
		"closes successfully": {
			ctx:     context.Background(),
			wantErr: false,
		},
		"returns an error for done context": {
			ctx:     doneCtx,
			wantErr: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			man := NewConsumerManager()
			err := man.Close(tc.ctx)
			assert.Equal(t, tc.wantErr, err != nil, "errors do not match")
		})
	}
}
