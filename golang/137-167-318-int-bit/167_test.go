package jzoffer

import (
	"reflect"
	"testing"
)

var _167TestCases = []struct {
	nums   []int
	target int
	res    []int
}{
	// 	* 输入：numbers = [2,7,11,15], target = 9
	//  * 输出：[1,2]
	//  * 解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
	//  *
	//  * 示例 2：
	//  *
	//  *
	//  * 输入：numbers = [2,3,4], target = 6
	//  * 输出：[1,3]
	//  * 解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。
	//  *
	//  * 示例 3：
	//  *
	//  *
	//  * 输入：numbers = [-1,0], target = -1
	//  * 输出：[1,2]
	//  * 解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		res:    []int{1, 2},
	},
	{
		nums:   []int{2, 3, 4},
		target: 6,
		res:    []int{1, 3},
	},
	{
		nums:   []int{-1, 0},
		target: -1,
		res:    []int{1, 2},
	},
}

func Test167(t *testing.T) {
	for _, tc := range _167TestCases {
		res := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("got %v, want %v", res, tc.res)
		}

		res = twoSumBinary(tc.nums, tc.target)
		if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("got %v, want %v", res, tc.res)
		}
	}
}
