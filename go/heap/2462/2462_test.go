package jzoffer

import "testing"

func Test2462(t *testing.T) {
	if got := totalCost([]int{48}, 1, 1); got != 48 {
		t.Errorf("totalCost() = %v, want %v", got, 48)
	}
	if got := totalCost([]int{18, 64, 12, 21, 21, 78, 36, 58, 88, 58, 99, 26, 92, 91, 53, 10, 24, 25, 20, 92, 73, 63, 51, 65, 87, 6, 17, 32, 14, 42, 46, 65, 43, 9, 75}, 13, 23); got != 223 {
		t.Errorf("totalCost() = %v, want %v", got, 223)
	}

	if got := totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4); got != 11 {
		t.Errorf("totalCost() = %v, want %v", got, 11)
	}
}
