package jzoffer

import (
	"reflect"
	"testing"
)

var _testcases = []struct {
	nums, expect []int
	k            int
}{
	{
		nums:   []int{1, 2, 3, 4, 5, 6, 7},
		k:      3,
		expect: []int{5, 6, 7, 1, 2, 3, 4},
	},
}

func Test189(t *testing.T) {
	for i, v := range _testcases {
		rotate(v.nums, v.k)
		if !reflect.DeepEqual(v.nums, v.expect) {
			t.Errorf("[%d]1.got:%v,expect:%v", i, v.nums, v.expect)
		}
	}
}
