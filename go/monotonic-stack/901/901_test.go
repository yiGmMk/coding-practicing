package jzoffer

import (
	"fmt"
	"testing"
)

func Test901(t *testing.T) {
	c := Constructor()
	for _, v := range []int{100, 80, 60, 70, 60, 75, 85} {
		fmt.Println(c.Next(v))
	}
}
