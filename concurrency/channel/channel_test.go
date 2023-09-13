package channel

import (
	"fmt"
	"runtime"
	"testing"
)

// ref https://go.dev/doc/effective_go#concurrency
// chan 能通过chan传递
// chan,函数都是"第一公民"
func TestChannelOfChannel(t *testing.T) {
	req := &Request{
		args:       []int{1, 2, 3},
		f:          sum,
		resultChan: make(chan int),
	}

	reqChan := make(chan *Request)
	quit := make(chan bool)
	go Serve(reqChan, quit)

	reqChan <- req
	fmt.Println("res:", <-req.resultChan)
	close(reqChan)
	close(req.resultChan)

	quit <- true
}

// Be sure not to confuse the ideas of
// concurrency structuring a program as independently executing components and     并发
// parallelism executing calculations in parallel for efficiency on multiple CPUs. 并行
// Although the concurrency features of Go can make some problems easy to structure as parallel computations,
// Go is a concurrent language, not a parallel one, and not all parallelization problems fit Go's model.
func TestParallelization(t *testing.T) {
	numCPU := runtime.NumCPU() * 2
	c := make(chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		go func(i int) {
			c <- i * i
		}(i)
	}
	for i := 0; i < numCPU; i++ {
		fmt.Println(<-c)
	}
	// all done
}
