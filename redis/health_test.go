package redis

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisPingCheck(t *testing.T) {
	type input struct {
		rbd     *redis.Client
		timeout time.Duration
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"returns an error if a database is nil": {input: input{}, wantErr: true},
		"returns an error if a context expired": {
			input: input{
				rbd:     redis.NewClient(&redis.Options{}),
				timeout: -1,
			},
			wantErr: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			check := PingCheck(tc.input.rbd, tc.input.timeout)
			err := check()
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
