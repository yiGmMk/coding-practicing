package jzoffer

import (
	"fmt"
	"testing"
)

func Test1466(t *testing.T) {
	// [[0,1],[1,3],[2,3],[4,0],[4,5]]
	fmt.Println(
		minReorder(6, [][]int{
			{0, 1}, {1, 3}, {2, 3},
			{4, 0}, {4, 5}}))
}
