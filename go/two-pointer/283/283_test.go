package jzoffer

import (
	"reflect"
	"testing"
)

var _testcase = []struct {
	nums, expect []int
}{
	{
		nums:   []int{1, 0, 1},
		expect: []int{1, 1, 0},
	},
	{
		nums:   []int{0, 1},
		expect: []int{1, 0},
	},

	{
		nums:   []int{0, 1, 0, 3, 12},
		expect: []int{1, 3, 12, 0, 0},
	},
}

func Test283(t *testing.T) {
	for i, v := range _testcase {
		data := make([]int, len(v.nums))
		copy(data, v.nums)
		moveZeroes1(data)
		if !reflect.DeepEqual(data, v.expect) {
			t.Errorf("[%d].got:%d,expect:%d", i, data, v.expect)
		}
	}
}
