package jzoffer

import (
	"fmt"
	"testing"
)

func TestGrayCode(t *testing.T) {
	n := 2
	format := fmt.Sprintf("%%0%db", n) // %0nb
	for _, v := range genGrayCode(2) {
		fmt.Printf(format+"\t", v)
	}
	fmt.Println()
}

var _testcase = []struct {
	n   int
	res []int
}{}
