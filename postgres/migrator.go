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
		return errors.Wrap(err, "acquire db conn")
	}
	defer stdlib.ReleaseConn(db, conn)

	migrator, err := migrate.NewMigrator(ctx, conn, _versionTable)
	if err != nil {
		return errors.Wrap(err, "create migrator")
	}

	if err = migrator.LoadMigrations(path); err != nil {
		return errors.Wrap(err, "load migrations")
	}

	if err = migrator.Migrate(ctx); err != nil {
		return errors.Wrap(err, "migrate db")
	}

	return nil
}
