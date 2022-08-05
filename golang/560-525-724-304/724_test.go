package jzoffer

import "testing"

var _724TestCases = []struct {
	nums []int
	res  int
}{
	//  * 示例 1：
	//  *
	//  *
	//  * 输入：nums = [1, 7, 3, 6, 5, 6]
	//  * 输出：3
	//  * 解释：
	//  * 中心下标是 3 。
	//  * 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
	//  * 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
	//  *
	//  *
	//  * 示例 2：
	//  *
	//  *
	//  * 输入：nums = [1, 2, 3]
	//  * 输出：-1
	//  * 解释：
	//  * 数组中不存在满足此条件的中心下标。
	//  *
	//  * 示例 3：
	//  *
	//  *
	//  * 输入：nums = [2, 1, -1]
	//  * 输出：0
	//  * 解释：
	//  * 中心下标是 0 。
	//  * 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
	//  * 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
	//  *
	{
		nums: []int{1, 7, 3, 6, 5, 6},
		res:  3,
	},
	{
		nums: []int{1, 2, 3},
		res:  -1,
	},
	{
		nums: []int{2, 1, -1},
		res:  0,
	},
}

func Test724(t *testing.T) {
	for _, tc := range _724TestCases {
		if res := pivotIndex(tc.nums); res != tc.res {
			t.Errorf("findPivot(%v) = %v, want %v", tc.nums, res, tc.res)
		}
	}
}
