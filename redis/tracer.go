package redis

import (
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

func RegisterTracing(rdb *redis.Client) {
	rdb.AddHook(redisotel.NewTracingHook())
}
