package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

// Migrate migrates the database from the provided directory path.
func Migrate(ctx context.Context, db *sql.DB, path string) error {
	var err error
	success := make(chan struct{}, 1)
	go func() {
		if err = goose.Up(db, path); err != nil {
			err = fmt.Errorf("up goose: %w", err)
		}
		success <- struct{}{}
	}()

	select {
	case <-success:
		return err
	case <-ctx.Done():
		return fmt.Errorf("context done: %w", ctx.Err())
	}
}
