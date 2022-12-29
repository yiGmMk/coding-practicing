package jzoffer

import "testing"

var _1292TestCase = []struct {
	mat    [][]int
	target int
	expect int
}{
	{
		mat:    [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
		target: 4,
		expect: 2,
	},
	{
		mat:    [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}},
		target: 1,
		expect: 0,
	},
}

func Test1482(t *testing.T) {
	for i, v := range _1292TestCase {
		got := maxSideLength1(v.mat, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = maxSideLength(v.mat, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
