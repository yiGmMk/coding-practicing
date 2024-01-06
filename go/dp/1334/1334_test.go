package jzoffer

import "testing"

func Test1334(t *testing.T) {
	// 5
	// [[0,1,2],[0,4,8],[1,2,3],[1,4,2],[2,3,1],[3,4,1]]
	// 2
	if got := findTheCity(5,
		[][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2); got != 0 {
		t.Errorf("findTheCity()=%v", got)
	}

	// [0,1,3],[1,2,1],[1,3,4],[2,3,1]
	if got := findTheCity(4,
		[][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4); got != 3 {
		t.Errorf("findTheCity()=%v", got)
	}
}
