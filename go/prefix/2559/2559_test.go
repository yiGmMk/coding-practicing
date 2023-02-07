package jzoffer

import (
	"fmt"
	"testing"
)

func Test2559(t *testing.T) {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))

	// [2,3,0]
}
