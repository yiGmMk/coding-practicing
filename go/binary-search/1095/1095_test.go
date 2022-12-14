package jzoffer

import "testing"

var _1095TestCases = []struct {
	nums   []int
	target int
	expect int
	max    int
}{
	{
		nums:   []int{1, 2, 3, 4, 5, 3, 1},
		target: 3,
		expect: 2,
		max:    5,
	},
	{
		nums:   []int{1, 3, 2},
		target: 1,
		expect: 0,
		max:    3,
	},
	{
		nums:   []int{0, 3, 1},
		target: 1,
		expect: 2,
		max:    3,
	},
	{
		nums:   []int{3, 11, 2},
		target: 11,
		expect: 1,
		max:    11,
	},
}

func Test1095(t *testing.T) {
	for i, v := range _1095TestCases {
		got := findInMountainArray(v.target, (*MountainArray)(&v.nums))
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = findInMountainArraySort(v.target, (*MountainArray)(&v.nums))
		if got != v.expect {
			t.Errorf("[%d]sort.got:%d,want:%d", i, got, v.expect)
		}
	}
}
