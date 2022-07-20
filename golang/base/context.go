package base

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	_trace_id = "trace_id"
)

func NewTraceID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func withTimeout(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var deal func(ctx context.Context)
	i := 0
	deal = func(ctx context.Context) {
		for {
			time.Sleep(time.Millisecond * 100)
			select {
			case <-ctx.Done():
				log.Println("timeout:", ctx.Err())
				return
			default:
				log.Println("deal:", i)
			}
			i++
		}
	}
	deal(ctx)
}

func wittCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Println("canceler:", ctx.Err())
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	log.Println("done")
	time.Sleep(time.Second * 2)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
