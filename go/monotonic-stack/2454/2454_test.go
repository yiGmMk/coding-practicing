package jzoffer

import (
	"fmt"
	"testing"
)

func Test2452(t *testing.T) {
	nums := []int{11, 13, 15, 12, 0, 15, 12, 11, 9}
	fmt.Println(secondGreaterElement(nums))
	// 15,15,-1,-1,12,-1,-1,-1,-1
}
