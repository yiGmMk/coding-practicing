package jzoffer

import (
	"testing"
)

var _209TestCase = []struct {
	nums        []int
	target, res int
}{
	{
		nums:   []int{2, 3, 1, 2, 4, 3},
		target: 7,
		res:    2,
	},
	{
		nums:   []int{1, 4, 4},
		target: 4,
		res:    1,
	},
	{
		nums:   []int{1, 1, 1, 1, 1},
		target: 11,
		res:    0,
	},
}

func Test209(t *testing.T) {
	for _, tc := range _209TestCase {
		res := minSubArrayLenTimeout(tc.target, tc.nums)
		if res != tc.res {
			t.Errorf("%v,target:%v,got:%v,res:%v", tc.nums, tc.target, res, tc.res)
		}
		res = minSubArrayLen(tc.target, tc.nums)
		if res != tc.res {
			t.Errorf("%v,target:%v,got:%v,res:%v", tc.nums, tc.target, res, tc.res)
		}
	}
}
