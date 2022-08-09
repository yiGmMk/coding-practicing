package redis

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	zeroRedis "github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	defaultDatabase = 0
	maxRetries      = 3
	idleConns       = 8
)

// NewRedisInmemory creates a new inmemory redis store
func NewRedisInmemory(t *testing.T, options ...zeroRedis.Option) *zeroRedis.Redis {
	r := miniredis.RunT(t)

	return zeroRedis.New(r.Addr(), options...)
}

func NewClient(t *testing.T, options ...zeroRedis.Option) *redis.Client {
	r := miniredis.RunT(t)

	return redis.NewClient(&redis.Options{
		Addr:         r.Addr(),
		DB:           defaultDatabase,
		MaxRetries:   maxRetries,
		MinIdleConns: idleConns,
	})
}
