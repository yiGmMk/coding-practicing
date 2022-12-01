package jzoffer

import (
	"testing"
)

var _137TestCases = []struct {
	nums []int
	res  int
}{
	{
		nums: []int{30000, 500, 100, 30000, 100, 30000, 100},
		res:  500,
	},
	{
		nums: []int{2, 2, 3, 2},
		res:  3,
	},
}

func Test137(t *testing.T) {
	for _, testCase := range _137TestCases {
		result := singleNumber(testCase.nums)
		if result != testCase.res {
			t.Errorf("%v = %d, expected %d", testCase.nums, result, testCase.res)
		}

		result = singleNumberOn(testCase.nums)
		if result != testCase.res {
			t.Errorf("%v = %d, expected %d", testCase.nums, result, testCase.res)
		}
	}
}
