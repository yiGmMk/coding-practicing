package jzoffer

import (
	"fmt"
	"testing"
	"time"
)

func Test22(t *testing.T) {
	now := time.Now()
	fmt.Println(generateParenthesis(2))
	fmt.Printf("time:%v\n", time.Since(now).Seconds())

	now = time.Now()
	fmt.Println(generateParenthesis(6))
	fmt.Printf("time:%v\n", time.Since(now).Seconds())
}
