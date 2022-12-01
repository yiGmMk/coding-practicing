package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	res := plusOne([]int{9})
	fmt.Println(res)
	res1 := make([]int, len(res)+1)
	res1[0] = 1
	fmt.Println(append(res1, res...))
}
