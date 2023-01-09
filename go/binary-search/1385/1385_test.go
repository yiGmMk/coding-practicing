package jzoffer

import "testing"

var _TestCase = []struct {
	nums   []int
	nums2  []int
	target int
	expect int
}{
	{
		nums:   []int{2, 1, 100, 3},
		nums2:  []int{-5, -2, 10, -3, 7},
		target: 6,
		expect: 1,
	},
	{
		nums:   []int{4, 5, 8},
		nums2:  []int{10, 9, 1, 8},
		target: 2,
		expect: 2,
	},
	{
		nums:   []int{1, 4, 2, 3},
		nums2:  []int{-4, -3, 6, 10, 20, 30},
		target: 3,
		expect: 2,
	},
}

func Test1300(t *testing.T) {
	for i, v := range _TestCase {
		got := findTheDistanceValue(v.nums, v.nums2, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = findTheDistanceValue1(v.nums, v.nums2, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
