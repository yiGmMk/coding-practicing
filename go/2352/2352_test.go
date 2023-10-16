package jzoffer

import "testing"

func Test2352(t *testing.T) {
	if 2 != equalPairs([][]int{{11, 1}, {1, 11}}) {
		t.Errorf("not equal")
	}
}
