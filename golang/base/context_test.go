package base

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestWithCance(t *testing.T) {
	wittCancel()
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	value := context.WithValue(ctx, _trace_id, NewTraceID())

	fmt.Println(
		"empty:", ctx.Value(_trace_id),
		"with value:", value.Value(_trace_id))

}

func TestContextTimeout(t *testing.T) {
	cur := time.Now()
	withTimeout(time.Second * 8)
	log.Println("finish,cost time:", time.Since(cur).Seconds())
}

func BenchmarkCommonParentCancel(b *testing.B) {
	root := context.WithValue(context.Background(), contextKey("key"), "value")
	shared, sharedcancel := context.WithCancel(root)
	defer sharedcancel()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		x := 0
		for pb.Next() {
			ctx, cancel := context.WithCancel(shared)
			if ctx.Value(contextKey("key")).(string) != "value" {
				b.Fatal("should not be reached")
			}
			for i := 0; i < 100; i++ {
				x /= x + 1
			}
			cancel()
			for i := 0; i < 100; i++ {
				x /= x + 1
			}
		}
	})
}
