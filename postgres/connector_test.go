package postgres

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	type (
		input struct {
			cfg  Config
			opts []Option
		}

		result struct {
			optApplied  bool
			maxOpenConn int
			wantErr     bool
		}
	)

	var optApplied bool
	tests := map[string]struct {
		input input
		want  result
	}{
		"does not connect to postgres with invalid url": {
			input: input{
				cfg:  Config{URL: "invalid-url"},
				opts: nil,
			},
			want: result{
				optApplied:  false,
				maxOpenConn: 0,
				wantErr:     true,
			},
		},
		"applies options and sets provided config successfully": {
			input: input{
				cfg: Config{
					URL:             "user=test host=test dbname=test",
					MaxOpenConn:     10,
					MaxConnLifetime: time.Second * 10,
					EnableTracing:   true,
				},
				opts: []Option{func(*pgx.ConnConfig) { optApplied = true }},
			},
			want: result{
				optApplied:  true,
				maxOpenConn: 10,
				wantErr:     false,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			optApplied = false // restore every time

			db, err := Connect(tc.input.cfg, tc.input.opts...)
			assert.Equal(t, tc.want.optApplied, optApplied)
			assert.Equal(t, tc.want.wantErr, err != nil, "errors do no match")

			if db != nil {
				assert.Equal(t, tc.want.maxOpenConn, db.Stats().MaxOpenConnections)
			}
		})
	}
}
