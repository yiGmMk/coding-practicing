package jzoffer

import "testing"

var _testcase = []struct {
	nums      []int
	k, expect int
}{
	{
		nums:   []int{9, 9, 9, 1, 2, 3},
		k:      3,
		expect: 12,
	},
	{
		nums:   []int{1, 5, 4, 2, 9, 9, 9},
		k:      3,
		expect: 15,
	},
}

func Test2461(t *testing.T) {
	for i, v := range _testcase {
		got := maximumSubarraySum(v.nums, v.k)
		if got != int64(v.expect) {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
