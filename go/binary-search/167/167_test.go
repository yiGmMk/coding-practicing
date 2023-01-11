package jzoffer

import (
	"testing"

	"github.com/samber/lo"
)

var _162TestCases = []struct {
	nums   []int
	target int
	expect []int
}{
	{
		nums:   []int{0, 0, 3, 4},
		target: 0,
		expect: []int{1, 2},
	},
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		expect: []int{1, 2},
	},
	{
		nums:   []int{2, 3, 4},
		target: 6,
		expect: []int{1, 3},
	},
	{
		nums:   []int{-1, 0},
		target: -1,
		expect: []int{1, 2},
	},
}

func Test167(t *testing.T) {
	for i, v := range _162TestCases {
		got := twoSumBinarySearch(v.nums, v.target)
		if !lo.Every(got, v.expect) {
			t.Errorf("[%d],got:%d,want:%d", i, got, v.expect)
		}

		got = twoSum(v.nums, v.target)
		if !lo.Every(got, v.expect) {
			t.Errorf("[%d],got:%d,want:%d", i, got, v.expect)
		}
	}
}
