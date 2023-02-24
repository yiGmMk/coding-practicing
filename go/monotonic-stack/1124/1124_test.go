package jzoffer

import (
	"testing"
)

var _testcase = []struct {
	hours  []int
	expect int
}{
	{
		hours:  []int{6, 6, 9},
		expect: 1,
	},
	{
		hours:  []int{9, 9, 6, 0, 6, 6, 9},
		expect: 3,
	},
	{
		hours:  []int{6, 6, 6},
		expect: 0,
	},
}

func Test1124(t *testing.T) {
	for i, v := range _testcase {
		hours := make([]int, len(v.hours))
		copy(hours, v.hours)
		got := longestWPI(hours)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}

		copy(hours, v.hours)
		got = longestWPI1(hours)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,expect:%d", i, got, v.expect)
		}
	}
}
