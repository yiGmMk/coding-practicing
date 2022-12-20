package jzoffer

import (
	"testing"
)

var _209TestCase = []struct {
	nums   []int
	target int
	expect int
}{
	{
		nums:   []int{2, 3, 1, 2, 4, 3},
		target: 7,
		expect: 2,
	},
	{
		nums:   []int{1, 4, 4},
		target: 4,
		expect: 1,
	},
	{
		nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		target: 11,
		expect: 0,
	},
}

func Test209(t *testing.T) {
	for i, v := range _209TestCase {
		got := minSubArrayLen(v.target, v.nums)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = minSubArrayLen1(v.target, v.nums)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,want:%d", i, got, v.expect)
		}
	}
}
