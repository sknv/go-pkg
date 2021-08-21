package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Config is a connection config.
type Config struct {
	URL             string
	MaxOpenConn     int           // maximum number of socket connections
	MaxConnLifetime time.Duration // connection age at which client closes the connection
	EnableTracing   bool
}

// Option configures *redis.Options.
type Option func(*redis.Options)

// Connect opens a redis connection.
func Connect(config Config, options ...Option) (*redis.Client, error) {
	opts, err := redis.ParseURL(config.URL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}

	// Apply config
	if config.MaxOpenConn != 0 {
		opts.PoolSize = config.MaxOpenConn
	}

	if config.MaxConnLifetime != 0 {
		opts.MaxConnAge = config.MaxConnLifetime
	}

	// Apply options
	for _, opt := range options {
		opt(opts)
	}

	rdb := redis.NewClient(opts)

	// Register tracing if needed
	if config.EnableTracing {
		RegisterTracing(rdb)
	}

	return rdb, nil
}
