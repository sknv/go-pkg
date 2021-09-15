package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
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

// MigrateEmbed migrates the database from the provided embed FS and directory path.
func MigrateEmbed(ctx context.Context, db *sql.DB, fs embed.FS, path string) error {
	var err error
	success := make(chan struct{}, 1)
	go func() {
		goose.SetBaseFS(fs)
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
