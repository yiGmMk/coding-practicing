package jzoffer

import "testing"

var _852TestCase = []struct {
	nums   []int
	target int
	expect int
}{
	{
		nums:   []int{0, 1, 0},
		expect: 1,
	},
	{
		nums:   []int{0, 2, 1, 0},
		expect: 1,
	},
	{
		nums:   []int{0, 10, 5, 2},
		expect: 1,
	},
	{
		nums:   []int{3, 4, 5, 1},
		expect: 2,
	},
	{
		nums:   []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
		expect: 2,
	},
}

func Test704(t *testing.T) {
	for i, v := range _852TestCase {
		got := peakIndexInMountainArray(v.nums)
		if got != v.expect {
			t.Errorf("[%d].got:%v,want:%v", i, got, v.expect)
		}
	}
}
