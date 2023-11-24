package jzoffer

import (
	"fmt"
	"testing"
)

func Test733(t *testing.T) {
	// [[1,1,1],[1,1,0],[1,0,1]]
	// 1
	// 1
	// 2
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}
