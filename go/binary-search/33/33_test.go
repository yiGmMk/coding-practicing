package jzoffer

import (
	"testing"
)

var _33TestCases = []struct {
	nums   []int
	target int
	expect int
	min    int
}{
	{
		nums:   []int{4, 5, 6, 7, 0, 1, 2},
		target: 0,
		expect: 4,
		min:    0,
	},
	{
		nums:   []int{1, 3},
		target: 1,
		expect: 0,
		min:    1,
	},
	{
		nums:   []int{3, 1},
		target: 1,
		expect: 1,
		min:    1,
	},
	{
		nums:   []int{3, 1, 2},
		target: 1,
		expect: 1,
		min:    1,
	},
}

func Test33(t *testing.T) {
	for i, v := range _33TestCases {
		min := findMin(v.nums)
		if min != v.min {
			t.Errorf("[%d].min,got:%d,want:%d", i, min, v.min)
		}

		got := search(v.nums, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
