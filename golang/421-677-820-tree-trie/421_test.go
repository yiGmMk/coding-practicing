package jzoffer

import (
	"fmt"
	"math/bits"
	"testing"
)

var _421TestCases = []struct {
	nums   []int
	expect int
}{
	{
		/*
		 * 输入：nums = [3,10,5,25,2,8]
		 * 输出：28
		 */
		nums:   []int{3, 10, 5, 25, 2, 8},
		expect: 28,
	},
	{
		/*
		 * 输入：nums = [0]
		 * 输出：0
		 *
		 */
		nums:   []int{0},
		expect: 0,
	},
	{
		/*
		 * 输入：nums = [2,4]
		 * 输出：6
		 */
		nums:   []int{2, 4},
		expect: 6,
	},
	{
		/* 输入：nums = [8,10,2]
		* 输出：10
		 */
		nums:   []int{8, 10, 2},
		expect: 10,
	},
	{ /* 输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
		 * 输出：127
		 */
		nums:   []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70},
		expect: 127,
	},
}

func Test421(t *testing.T) {
	for i, testCase := range _421TestCases {
		fmt.Println(bits.Len(uint(testCase.expect)))
		actual := findMaximumXOR(testCase.nums)
		if actual != testCase.expect {
			t.Errorf("%d: expect: %v,  	actual: %v\n", i, testCase.expect, actual)
		}
	}
}
