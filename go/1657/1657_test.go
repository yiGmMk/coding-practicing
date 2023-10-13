package jzoffer

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
)

func TestEq(t *testing.T) {
	fmt.Println(lo.Slice[int]([]int{1, 2, 3, 4, 5}, 1, 3))
}
