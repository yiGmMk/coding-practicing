package jzoffer

import "testing"

var _1818TestCase = []struct {
	nums1, nums2 []int
	expect       int
}{
	{
		nums1:  []int{1, 28, 21},
		nums2:  []int{9, 21, 20},
		expect: 9,
	},
	{
		nums1:  []int{1, 7, 5},
		nums2:  []int{2, 3, 5},
		expect: 3,
	},
	{
		nums1:  []int{2, 4, 6, 8, 10},
		nums2:  []int{2, 4, 6, 8, 10},
		expect: 0,
	},
	{
		nums1:  []int{1, 10, 4, 4, 2, 7},
		nums2:  []int{9, 3, 5, 1, 7, 4},
		expect: 20,
	},
}

func Test1482(t *testing.T) {
	for i, v := range _1818TestCase {
		got := minAbsoluteSumDiff(v.nums1, v.nums2)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = minAbsoluteSumDiff1(v.nums1, v.nums2)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,want:%d", i, got, v.expect)
		}
	}
}
