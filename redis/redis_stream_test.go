package redis

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestStream(t *testing.T) {

	client := NewClient(t)

	topic := "redis_stream"
	ctx := context.Background()

	err := client.XAdd(ctx, &redis.XAddArgs{Stream: topic}).Err()

	if err != nil {
		t.Error(err)
	}
	// client.XGroupCreate()
}
