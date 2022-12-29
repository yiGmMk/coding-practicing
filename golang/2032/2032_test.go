package jzoffer

import (
	"testing"

	"github.com/samber/lo"
)

var _1292TestCase = []struct {
	nums1, nums2, nums3 []int
	expect              []int
}{
	{
		nums1:  []int{1, 2, 1, 3},
		nums2:  []int{2, 3},
		nums3:  []int{3},
		expect: []int{3, 2},
	},
	{
		nums1:  []int{1, 3},
		nums2:  []int{2, 3},
		nums3:  []int{1, 2},
		expect: []int{3, 2, 1},
	},
	{
		nums1:  []int{1, 2, 2},
		nums2:  []int{4, 3, 3},
		nums3:  []int{5},
		expect: []int{},
	},
}

func Test1482(t *testing.T) {
	for i, v := range _1292TestCase {
		got := twoOutOfThree(v.nums1, v.nums2, v.nums3)
		if l, r := lo.Difference(got, v.expect); len(l) > 0 || len(r) > 0 {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = twoOutOfThree1(v.nums1, v.nums2, v.nums3)
		if l, r := lo.Difference(got, v.expect); len(l) > 0 || len(r) > 0 {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = twoOutOfThree2(v.nums1, v.nums2, v.nums3)
		if l, r := lo.Difference(got, v.expect); len(l) > 0 || len(r) > 0 {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = twoOutOfThree3(v.nums1, v.nums2, v.nums3)
		if l, r := lo.Difference(got, v.expect); len(l) > 0 || len(r) > 0 {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
