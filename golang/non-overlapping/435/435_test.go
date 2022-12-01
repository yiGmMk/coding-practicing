package jzoffer

import "testing"

var _435TestCase = []struct {
	intervals [][]int
	expect    int
}{
	{
		intervals: [][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}, {5, 8}},
		expect:    3,
	},
	{
		intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		expect:    1,
	},
	{
		intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
		expect:    2,
	},
	{
		intervals: [][]int{{1, 2}, {2, 3}},
		expect:    0,
	},
}

func Test435(t *testing.T) {
	for _, v := range _435TestCase {
		res := eraseOverlapIntervals(v.intervals)
		if res != v.expect {
			t.Errorf("1.want:%d,got:%d", v.expect, res)
		}

		res = eraseOverlapIntervalsTimeOut(v.intervals)
		if res != v.expect {
			t.Errorf("2.want:%d,got:%d", v.expect, res)
		}
	}
}
