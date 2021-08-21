package redis

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/sknv/go-pkg/log"
)

type Logger struct{}

func (Logger) BeforeProcess(ctx context.Context, _ redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (Logger) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	log.Extract(ctx).WithField("op", "redis").Debug(cmd) // always use debug level
	return nil
}

func (Logger) BeforeProcessPipeline(ctx context.Context, _ []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (Logger) AfterProcessPipeline(context.Context, []redis.Cmder) error {
	return nil
}

// RegisterLogger registers a redis logger hook.
func RegisterLogger(rdb *redis.Client) {
	rdb.AddHook(Logger{})
}
