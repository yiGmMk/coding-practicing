package redis

import (
	"context"
	"log"
	"testing"
)

func TestListBase(t *testing.T) {

	r := NewRedisInmemory(t)
	log.Println(r.Addr)

	ctx := context.Background()
	count, err := r.RpushCtx(ctx, "list", "a", "b", "c")
	if err != nil {
		t.Error(err)
	}
	if count != 3 {
		t.Errorf("count should be 3, but got %d", count)
	}

	val, err := r.LlenCtx(ctx, "list")
	if err != nil {
		t.Error(err)
	}
	if val != 3 {
		t.Errorf("count should be 3, but got %d", val)
	}

	popVal, err := r.LpopCtx(ctx, "list")
	if err != nil {
		t.Error(err)
	}

	if popVal != "a" {
		t.Errorf("popVal should be a, but got %s", popVal)
	}
}
