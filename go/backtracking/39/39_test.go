package jzoffer

import (
	"fmt"
	"testing"
)

func Test39(t *testing.T) {
	var (
		a = [4]int{1, 2, 3}
		b = [4]int{1, 2, 3, 4}
		c = [4]int{1, 2, 3}
	)
	fmt.Println(a, b, c, a == b, a == c, b == c)

	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
