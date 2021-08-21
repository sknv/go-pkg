package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

const (
	DriverName = "pgx"
)

// Config is a connection config.
type Config struct {
	URL             string
	MaxOpenConn     int           // maximum number of open connections
	MaxConnLifetime time.Duration // maximum amount of time a connection may be reused
	EnableTracing   bool
}

// Option configures *pgx.ConnConfig.
type Option func(*pgx.ConnConfig)

// Connect opens a db connection.
func Connect(config Config, options ...Option) (*sql.DB, error) {
	connConfig, err := pgx.ParseConfig(config.URL)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	// Register tracing if needed
	driverName := DriverName
	if config.EnableTracing {
		driverName, err = RegisterTracer(driverName)
		if err != nil {
			return nil, fmt.Errorf("register tracer: %w", err)
		}
	}

	// Apply options
	for _, opt := range options {
		opt(connConfig)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	// Apply config
	if config.MaxOpenConn != 0 {
		db.SetMaxOpenConns(config.MaxOpenConn)
		db.SetMaxIdleConns(config.MaxOpenConn)
	}

	if config.MaxConnLifetime != 0 {
		db.SetConnMaxLifetime(config.MaxConnLifetime)
	}

	return db, nil
}
