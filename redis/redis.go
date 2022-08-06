package redis

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// NewRedisInmemory creates a new inmemory redis store
func NewRedisInmemory(t *testing.T, options ...redis.Option) *redis.Redis {
	r := miniredis.RunT(t)

	return redis.New(r.Addr(), options...)
}
