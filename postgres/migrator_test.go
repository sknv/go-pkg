package postgres

import (
	"context"
	"database/sql"
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type input struct {
		ctx context.Context
		db  *sql.DB
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"returns an error if context is done": {
			input:   input{ctx: canceledCtx},
			wantErr: true,
		},
		"returns an error if db is nil": {
			input:   input{ctx: context.Background()},
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := Migrate(tc.input.ctx, tc.input.db, "any")
			assert.Equal(t, tc.wantErr, err != nil)

			err = MigrateEmbed(tc.input.ctx, tc.input.db, embed.FS{}, "any")
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
