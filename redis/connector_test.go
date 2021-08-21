package redis

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	type (
		input struct {
			cfg  Config
			opts []Option
		}

		result struct {
			optApplied      bool
			maxOpenConn     int
			maxConnLifetime time.Duration
			wantErr         bool
		}
	)

	var optApplied bool
	tests := map[string]struct {
		input input
		want  result
	}{
		"does not connect to redis with invalid url": {
			input: input{
				cfg:  Config{URL: "invalid-url"},
				opts: nil,
			},
			want: result{
				optApplied:      false,
				maxOpenConn:     0,
				maxConnLifetime: 0,
				wantErr:         true,
			},
		},
		"applies options and sets provided config successfully": {
			input: input{
				cfg: Config{
					URL:             "redis://localhost:6379",
					MaxOpenConn:     10,
					MaxConnLifetime: time.Second * 10,
				},
				opts: []Option{func(*redis.Options) { optApplied = true }},
			},
			want: result{
				optApplied:      true,
				maxOpenConn:     10,
				maxConnLifetime: time.Second * 10,
				wantErr:         false,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			optApplied = false // restore every time

			rdb, err := Connect(tc.input.cfg, tc.input.opts...)
			assert.Equal(t, tc.want.optApplied, optApplied)
			assert.Equal(t, tc.want.wantErr, err != nil, "errors do no match")

			if rdb != nil {
				assert.Equal(t, tc.want.maxOpenConn, rdb.Options().PoolSize)
				assert.Equal(t, tc.want.maxConnLifetime, rdb.Options().MaxConnAge)
			}
		})
	}
}
