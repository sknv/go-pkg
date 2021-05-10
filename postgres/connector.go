package postgres

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
)

const DriverName = "pgx"

type Config struct {
	URI             string        `mapstructure:"uri"`
	MaxOpenConn     int           `mapstructure:"max_open_conn"`
	MaxIdleConn     int           `mapstructure:"max_idle_conn"`
	MaxConnLifetime time.Duration `mapstructure:"max_conn_lifetime"`
}

// Option configures *pgx.ConnConfig.
type Option func(*pgx.ConnConfig)

func Connect(config Config, options ...Option) (*sql.DB, error) {
	connConfig, err := pgx.ParseConfig(config.URI)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ParseConfig")
	}

	// Apply options
	for _, opt := range options {
		opt(connConfig)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)
	db, err := sql.Open(DriverName, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open")
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
