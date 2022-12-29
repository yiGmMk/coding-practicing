package jzoffer

import "testing"

var _1818TestCase = []struct {
	nums1  []int
	target int
	expect int
}{
	{
		nums1:  []int{2, 3, 3, 4, 6, 7},
		target: 12,
		expect: 61,
	},
	{
		nums1:  []int{3, 5, 6, 7},
		target: 9,
		expect: 4,
	},
	{
		nums1:  []int{3, 3, 6, 8},
		target: 10,
		expect: 6,
	},
}

func Test1482(t *testing.T) {
	for i, v := range _1818TestCase {
		got := numSubseq(v.nums1, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
