package redis

import (
	"testing"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestLock(t *testing.T) {
	server := NewRedisInmemory(t)

	lock := redis.NewRedisLock(server, "test")
	got, err := lock.Acquire()
	if err != nil || !got {
		t.Errorf("should be true,%+v,%+v", got, err)
	}
}
