package jzoffer

import "testing"

var _testcase = []struct {
	grid   [][]int
	expect int
}{

	{
		grid:   [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		expect: 9,
	},
}

func Test1139(t *testing.T) {
	for i, v := range _testcase {
		got := largest1BorderedSquare(v.grid)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
