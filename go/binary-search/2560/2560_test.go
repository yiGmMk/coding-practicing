package jzoffer

import "testing"

var _testcase = []struct {
	nums      []int
	k, expect int
}{
	{
		nums:   []int{2, 3, 5, 9},
		k:      2,
		expect: 5,
	},
	{
		nums:   []int{2, 7, 9, 3, 1},
		expect: 2,
		k:      2,
	},
}

func Test2560(t *testing.T) {

	for i, v := range _testcase {
		got := minCapability(v.nums, v.k)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
