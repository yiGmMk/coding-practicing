package jzoffer

import "testing"

var _713TestCases = []struct {
	nums []int
	k    int
	res  int
}{
	/*
	 * 示例 1：
	 *
	 *
	 * 输入：nums = [10,5,2,6], k = 100
	 * 输出：8
	 * 解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
	 * 需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
	 *
	 *
	 * 示例 2：
	 *
	 *
	 * 输入：nums = [1,2,3], k = 0
	 * 输出：0
	 */
	{
		nums: []int{10, 5, 2, 6},
		k:    100,
		res:  8,
	},
	{
		nums: []int{1, 2, 3},
		k:    0,
		res:  0,
	},
}

func Test713(t *testing.T) {
	for _, tc := range _713TestCases {
		if res := numSubarrayProductLessThanK(tc.nums, tc.k); res != tc.res {
			t.Errorf("nums: %v, k: %v, res: %v,expect:%v", tc.nums, tc.k, res, tc.res)
		}

		if res := numSubarrayProductLessThanKBinary(tc.nums, tc.k); res != tc.res {
			t.Errorf("nums: %v, k: %v, res: %v,expect:%v", tc.nums, tc.k, res, tc.res)
		}

		if res := numSubarrayProductLessThanK2(tc.nums, tc.k); res != tc.res {
			t.Errorf("nums: %v, k: %v, res: %v,expect:%v", tc.nums, tc.k, res, tc.res)
		}
	}
}
