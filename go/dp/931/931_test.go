package jzoffer

import (
	"testing"
)

func Test731(t *testing.T) {
	if got := minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}); got != 13 {
		t.Errorf("minFallingPathSum() = %v, want %v", got, 12)
	}
}
