package jzoffer

import "testing"

var _testcase = []struct {
	ranges    []int
	n, expect int
}{
	{
		ranges: []int{1, 2, 1, 0, 2, 1, 0, 1},
		n:      7,
		expect: 3,
	},
	{
		ranges: []int{3, 4, 1, 1, 0, 0},
		n:      5,
		expect: 1,
	},
}

func Test1326(t *testing.T) {
	for i, v := range _testcase {
		got := minTapsGreedy(v.n, v.ranges)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
