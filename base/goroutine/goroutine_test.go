package base

import (
	"fmt"
	"testing"
	"time"
)

// 通过channel返回gotoutine的返回状态
func TestExit(t *testing.T) {
	c := make(chan error, 1)
	go func() {
		time.Sleep(time.Millisecond * 100)

		c <- nil
	}()
	err := <-c
	fmt.Printf("goroutine exit,err:%v", err)
}
