package jzoffer

import "testing"

var _1803TestCase = []struct {
	nums      []int
	low, high int
	expect    int
}{
	{
		nums:   []int{1, 4, 2, 7},
		low:    2,
		high:   6,
		expect: 6,
	},
	{
		nums:   []int{9, 8, 4, 2, 1},
		low:    5,
		high:   14,
		expect: 8,
	},
}

func Test1803(t *testing.T) {
	for i, v := range _1803TestCase {
		got := countPairsN2(v.nums, v.low, v.high)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = countPairs(v.nums, v.low, v.high)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
