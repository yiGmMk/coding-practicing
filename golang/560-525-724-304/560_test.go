package jzoffer

import "testing"

var _560TestCases = []struct {
	nums []int
	k    int
	res  int
}{
	/*
	 * 示例 1：
	 *
	 *
	 * 输入：nums = [1,1,1], k = 2
	 * 输出：2
	 *
	 *
	 * 示例 2：
	 *
	 *
	 * 输入：nums = [1,2,3], k = 3
	 * 输出：2
	 */
	{
		nums: []int{1, -1, 0},
		k:    0,
		res:  3,
	},
	{
		nums: []int{1},
		k:    0,
		res:  0,
	},
	{
		nums: []int{1, 2, 3},
		k:    3,
		res:  2,
	},
	{
		nums: []int{1, 1, 1},
		k:    2,
		res:  2,
	},
}

func Test560(t *testing.T) {
	for _, tc := range _560TestCases {
		if res := subarraySum(tc.nums, tc.k); res != tc.res {
			t.Errorf("nums: %v, k: %v, res: %v,expect:%v", tc.nums, tc.k, res, tc.res)
		}
		if res := subarraySumHash(tc.nums, tc.k); res != tc.res {
			t.Errorf("nums: %v, k: %v, res: %v,expect:%v", tc.nums, tc.k, res, tc.res)
		}
	}
}
