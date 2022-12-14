package jzoffer

import (
	"testing"

	"github.com/samber/lo"
)

var _162TestCases = []struct {
	nums   []int
	expect []int
}{
	{
		nums:   []int{3, 1, 2},
		expect: []int{3},
	},
	{
		nums:   []int{1, 2, 1, 3, 5, 6, 4},
		expect: []int{2, 6},
	},
	{
		nums:   []int{4, 5, 6, 7, 0, 1, 2},
		expect: []int{7, 2},
	},
	{
		nums:   []int{1, 3},
		expect: []int{3},
	},
	{
		nums:   []int{3, 1},
		expect: []int{3},
	},
}

func Test33(t *testing.T) {
	for i, v := range _162TestCases {
		maxId := findPeakElement(v.nums)
		if maxId < 0 || maxId >= len(v.nums) || lo.Count(v.expect, v.nums[maxId]) <= 0 {
			t.Errorf("[%d].max,got:%d,want:%d", i, maxId, v.expect)
		}
	}
}
