package postgres

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

const DriverName = "pgx"

type Config struct {
	URI             string        `mapstructure:"uri"`
	MaxOpenConn     int           `mapstructure:"max_open_conn"`
	MaxIdleConn     int           `mapstructure:"max_idle_conn"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
	MigrationPath   string        `mapstructure:"migration_path"`
}

func (c *Config) Valid() error {
	var err error
	if c.URI == "" {
		err = multierr.Append(err, errors.New("empty uri"))
	}
	if c.MigrationPath == "" {
		err = multierr.Append(err, errors.New("empty migration path"))
	}
	return err
}

// Option configures *pgx.ConnConfig.
type Option func(*pgx.ConnConfig)

func Connect(config Config, options ...Option) (*sql.DB, error) {
	connConfig, err := pgx.ParseConfig(config.URI)
	if err != nil {
		return nil, errors.Wrap(err, "parse config")
	}

	// Apply options
	for _, opt := range options {
		opt(connConfig)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)
	db, err := sql.Open(DriverName, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "open db")
	}

	// Apply config
	if config.MaxOpenConn != 0 {
		db.SetMaxOpenConns(config.MaxOpenConn)
	}

	if config.MaxIdleConn != 0 {
		db.SetMaxIdleConns(config.MaxIdleConn)
	}

	if config.MaxConnLifetime != 0 {
		db.SetConnMaxLifetime(config.MaxConnLifetime)
	}

	return db, nil
}
