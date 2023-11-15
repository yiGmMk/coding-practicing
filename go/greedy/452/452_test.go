package jzoffer

import "testing"

func Test452(t *testing.T) {
	tcs := []struct {
		ps  [][]int
		res int
	}{
		{
			ps:  [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			res: 2,
		},
	}

	for _, v := range tcs {
		if got := findMinArrowShots(v.ps); got != v.res {
			t.Errorf("got:%v,want:%d", got, v.res)
		}
	}
}
