package concurrency

import (
	"fmt"
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	gen := func(n int) chan int {
		out := make(chan int)
		go func() {
			for i := 0; i < 100; i++ {
				out <- i * n
			}
			close(out)
		}()
		return out
	}
	in1 := gen(10)
	in2 := gen(1)
	in3 := gen(100)
	out := Merge(in1, in2, in3)
	val := 0
	for v := range out {
		fmt.Println(v)
		val++
	}
	log.Println("num:", val)
}
