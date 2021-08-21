package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/heptiolabs/healthcheck"
)

// PingCheck returns a healthcheck.Check that validates connectivity to a redis client using Ping().
func PingCheck(rdb *redis.Client, timeout time.Duration) healthcheck.Check {
	return func() error {
		if rdb == nil {
			return errors.New("database is nil")
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		return rdb.Ping(ctx).Err()
	}
}
