package jzoffer

import "testing"

var _testCase = []struct {
	nums   []int
	expect int
}{
	{
		nums:   []int{1, 3, 4, 4, 2, 3, 2, 2, 1, 1, 4, 1, 3, 2, 2, 5, 5, 2, 3, 1, 2, 3, 1, 3, 3, 1, 2, 3, 3, 5, 1},
		expect: 2,
	},
	{
		nums:   []int{1, 1, 1},
		expect: 3,
	},
	{
		nums:   []int{1, 2, 3},
		expect: 0,
	},
	{
		nums:   []int{2, 1, 6, 4},
		expect: 1,
	},
}

func Test1664(t *testing.T) {
	for i, v := range _testCase {
		got := waysToMakeFair(v.nums)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}

		got = waysToMakeFair1(v.nums)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,expect:%d", i, got, v.expect)
		}

		got = waysToMakeFair2(v.nums)
		if got != v.expect {
			t.Errorf("[%d]2.got:%d,expect:%d", i, got, v.expect)
		}
	}
}
