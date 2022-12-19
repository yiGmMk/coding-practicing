package jzoffer

import "testing"

var _1283TestCase = []struct {
	nums      []int
	threshold int
	expect    int
}{
	{
		nums:      []int{1, 2, 5, 9},
		threshold: 6,
		expect:    5,
	},
	{
		nums:      []int{2, 3, 5, 7, 11},
		threshold: 11,
		expect:    3,
	},
	{
		nums:      []int{19},
		threshold: 5,
		expect:    4,
	},
}

func Test1283(t *testing.T) {
	for i, v := range _1283TestCase {
		got := smallestDivisor(v.nums, v.threshold)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
