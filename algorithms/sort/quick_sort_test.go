package algorithms

import (
	"reflect"
	"sort"
	"testing"
)

var _sortTestCases = []SortTestCase{
	{
		nums:   []int{1, 9, 10, 100, 1000, 2, 3, 8, 1000_000, 4, 5, 6, 7},
		sorted: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 1000_000},
	},
	{
		nums:   []int{1, 5, -1},
		sorted: []int{-1, 1, 5},
	},
	{
		nums:   []int{0},
		sorted: []int{0},
	},
}

func checkSort(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	for _, v := range _sortTestCases {
		QuickSort(v.nums)
		if !reflect.DeepEqual(v.nums, v.sorted) ||
			!checkSort(v.nums) ||
			!sort.IntsAreSorted(v.nums) {
			t.Errorf("QuickSort failed,res:%v", v.nums)
		}
	}
}
