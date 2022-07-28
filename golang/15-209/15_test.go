package jzoffer

import (
	"reflect"
	"testing"
)

var _15TestCase = []struct {
	nums []int
	res  [][]int
}{
	{
		nums: []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
		res: [][]int{
			{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0},
		},
	},
	{
		nums: []int{0, 0, 0},
		res:  [][]int{{0, 0, 0}},
	},
	{
		nums: []int{1, -1, -1, 0},
		res:  [][]int{{-1, 0, 1}},
	},
	/*
	 * 示例 1：
	 *
	 *
	 * 输入：nums = [-1,0,1,2,-1,-4]
	 * 输出：[[-1,-1,2],[-1,0,1]]
	 *
	 *
	 * 示例 2：
	 *
	 *
	 * 输入：nums = []
	 * 输出：[]
	 *
	 *
	 * 示例 3：
	 *
	 *
	 * 输入：nums = [0]
	 * 输出：[]
	 */
	{
		nums: []int{-1, 0, 1, 2, -1, -4},
		res: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
	},
	{
		nums: []int{},
		res:  [][]int{},
	},
	{
		nums: []int{0},
		res:  [][]int{},
	},
}

func Test15(t *testing.T) {

	a, b := [][]int{{1, 2}}, [][]int{{1, 2}}
	if !reflect.DeepEqual(a, b) {
		t.Error("error")
	}

	for _, v := range _15TestCase {
		res := threeSum(v.nums)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("failed! Got %v, Want %v", res, v.res)
		}

		res = threeSum2(v.nums)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("failed! Got %v, Want %v", res, v.res)
		}
	}
}
