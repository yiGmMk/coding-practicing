package jzoffer

import "testing"

var _testcase = []struct {
	nums      []int
	k, expect int
}{
	{
		nums:   []int{7, 7, 7, 7},
		k:      2,
		expect: 0,
	},
	{
		nums:   []int{13, 5, 1, 8, 21, 2},
		k:      3,
		expect: 8,
	},
	{
		nums:   []int{1, 3, 1},
		k:      2,
		expect: 2,
	},
}

func Test2517(t *testing.T) {
	for i, v := range _testcase {
		got := maximumTastiness(v.nums, v.k)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
