package redis

import (
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestLock(t *testing.T) {
	r := miniredis.RunT(t)

	fmt.Println(r.Addr())

	server := redis.New(r.Addr())
	lock := redis.NewRedisLock(server, "test")
	got, err := lock.Acquire()
	if err != nil || !got {
		t.Errorf("should be true,%+v,%+v", got, err)
	}
}
