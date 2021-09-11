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
			opts []func(*bool) Option
		}

		result struct {
			optApplied      bool
			maxOpenConn     int
			maxConnLifetime time.Duration
			wantErr         bool
		}
	)

	tests := map[string]struct {
		input      input
		optApplied bool
		want       result
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
				opts: []func(*bool) Option{
					func(opt *bool) Option {
						return func(*redis.Options) { *opt = true }
					},
				},
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
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			opts := make([]Option, 0, len(tc.input.opts))
			for _, opt := range tc.input.opts {
				opts = append(opts, opt(&tc.optApplied))
			}
			rdb, err := Connect(tc.input.cfg, opts...)
			assert.Equal(t, tc.want.optApplied, tc.optApplied)
			assert.Equal(t, tc.want.wantErr, err != nil, "errors do no match")

			if rdb != nil {
				assert.Equal(t, tc.want.maxOpenConn, rdb.Options().PoolSize)
				assert.Equal(t, tc.want.maxConnLifetime, rdb.Options().MaxConnAge)
			}
		})
	}
}
