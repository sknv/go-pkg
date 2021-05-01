package postgres

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/tern/migrate"
	"github.com/pkg/errors"
)

const _versionTable = "schema_version"

func Migrate(ctx context.Context, db *sql.DB, path string) error {
	conn, err := stdlib.AcquireConn(db)
	if err != nil {
		return errors.Wrap(err, "stdlib.AcquireConn")
	}
	defer stdlib.ReleaseConn(db, conn)

	migrator, err := migrate.NewMigrator(ctx, conn, _versionTable)
	if err != nil {
		return errors.Wrap(err, "migrate.NewMigrator")
	}

	if err = migrator.LoadMigrations(path); err != nil {
		return errors.Wrap(err, "migrate.LoadMigrations")
	}

	if err = migrator.Migrate(ctx); err != nil {
		return errors.Wrap(err, "migrate.Migrate")
	}

	return nil
}
