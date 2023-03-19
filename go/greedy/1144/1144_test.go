package jzoffer

import (
	"testing"
)

var _testcase = []struct {
	nums   []int
	expect int
}{
	{
		nums:   []int{2, 1, 2, 1},
		expect: 0,
	},
	{
		nums:   []int{7, 4, 8, 9, 7, 7, 5},
		expect: 6,
	},
	{
		nums:   []int{2, 7, 10, 9, 8, 9},
		expect: 4,
	},
	{
		nums:   []int{2, 1, 2},
		expect: 0,
	},
	{
		nums:   []int{1, 2, 3},
		expect: 2,
	},
}

func Test1144(t *testing.T) {
	for i, v := range _testcase {
		got := movesToMakeZigzag(v.nums)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}

		got = movesToMakeZigzag0(v.nums)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}

}
