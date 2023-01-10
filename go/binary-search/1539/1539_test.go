package jzoffer

import (
	"testing"
)

var _1539TestCase = []struct {
	arr    []int
	k      int
	expect int
}{
	{
		arr:    []int{5, 6, 7, 8, 9},
		k:      9,
		expect: 14,
	},
	{
		arr:    []int{2},
		k:      1,
		expect: 1,
	},
	{
		arr:    []int{2, 3, 4, 7, 11},
		k:      5,
		expect: 9,
	},
	{
		arr:    []int{1, 2, 3, 4},
		k:      2,
		expect: 6,
	},
}

func Test1539(t *testing.T) {
	for i, v := range _1539TestCase {
		got := findKthPositiveOk(v.arr, v.k)
		if got != v.expect {
			t.Errorf("[%d].got:%v,wangt:%v", i, got, v.expect)
		}
		got = findKthPositiveOn1(v.arr, v.k)
		if got != v.expect {
			t.Errorf("[%d]0.1.got:%v,wangt:%v", i, got, v.expect)
		}

		got = findKthPositive(v.arr, v.k)
		if got != v.expect {
			t.Errorf("[%d]1.got:%v,wangt:%v", i, got, v.expect)
		}

		got = findKthPositive1(v.arr, v.k)
		if got != v.expect {
			t.Errorf("[%d]1.1.got:%v,wangt:%v", i, got, v.expect)
		}
	}
}
