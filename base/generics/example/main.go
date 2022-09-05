package main

import (
	"fmt"
	"sort"

	"github.com/duke-git/lancet/v2/lancetconstraints"
)

func Sort[T lancetconstraints.Number](items []T) {
	sort.Slice(items, func(x, y int) bool {
		return items[x] < items[y]
	})
	return
}

func Max[T lancetconstraints.Number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func main() {
	a, b := []int{1, 4, 2}, []float64{1, 4, 2}
	fmt.Println(a, b)
	Sort(a)
	Sort(b)
	fmt.Println(a, b, Max(1, 2), Max(1.1, 2.2))
}
