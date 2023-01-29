package jzoffer

import (
	"reflect"
	"testing"
)

var _testcase = []struct {
	nums []int
	res  []int
}{
	{
		nums: []int{-1},
		res:  []int{1},
	},
	{
		nums: []int{-4, -1, 0, 3, 10},
		res:  []int{0, 1, 9, 16, 100},
	},
	{
		nums: []int{-7, -3, 2, 3, 11},
		res:  []int{4, 9, 9, 49, 121},
	},
}

func Test977(t *testing.T) {
	for i, v := range _testcase {
		nums := make([]int, len(v.nums))
		copy(nums, v.nums)

		got := sortedSquares(nums)
		if !reflect.DeepEqual(got, v.res) {
			t.Errorf("[%d].got:%v,expect:%v", i, got, v.res)
		}

		nums = make([]int, len(v.nums))
		copy(nums, v.nums)
		got = sortedSquares1(nums)
		if !reflect.DeepEqual(got, v.res) {
			t.Errorf("[%d]1.got:%v,expect:%v", i, got, v.res)
		}

		nums = make([]int, len(v.nums))
		copy(nums, v.nums)
		got = sortedSquares3(nums)
		if !reflect.DeepEqual(got, v.res) {
			t.Errorf("[%d]3.got:%v,expect:%v", i, got, v.res)
		}

		nums = make([]int, len(v.nums))
		copy(nums, v.nums)
		got = sortedSquares2(nums)
		if !reflect.DeepEqual(got, v.res) {
			t.Errorf("[%d]2.got:%v,expect:%v", i, got, v.res)
		}
	}
}
