package jzoffer

import "testing"

var _1482TestCase = []struct {
	nums   []int
	m, k   int
	expect int
}{
	{
		nums:   []int{1, 10, 3, 10, 2},
		m:      3,
		k:      1,
		expect: 3,
	},
	{
		nums:   []int{1, 10, 3, 10, 2},
		m:      3,
		k:      2,
		expect: -1,
	},
	{
		nums:   []int{7, 7, 7, 7, 12, 7, 7},
		m:      2,
		k:      3,
		expect: 12,
	},
	{
		nums:   []int{1000000000, 1000000000},
		m:      1,
		k:      1,
		expect: 1000000000,
	},
}

func Test1482(t *testing.T) {
	for i, v := range _1482TestCase {
		got := minDays(v.nums, v.m, v.k)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = minDays1(v.nums, v.m, v.k)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,want:%d", i, got, v.expect)
		}

		got = minDays2(v.nums, v.m, v.k)
		if got != v.expect {
			t.Errorf("[%d]2.got:%d,want:%d", i, got, v.expect)
		}
	}
}
