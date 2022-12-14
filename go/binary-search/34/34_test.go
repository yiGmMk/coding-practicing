package jzoffer

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var _34TestCase = []struct {
	in     []int
	target int
	expect []int
}{
	{
		in:     []int{1},
		target: 1,
		expect: []int{0, 0},
	},
	{
		in:     []int{5, 7, 7, 8, 8, 10},
		target: 8,
		expect: []int{3, 4},
	},
	{
		in:     []int{5, 7, 7, 8, 8, 10},
		target: 6,
		expect: []int{-1, -1},
	},
	{
		in:     []int{},
		target: 0,
		expect: []int{-1, -1},
	},
}

func Test34(t *testing.T) {
	fmt.Println(sort.SearchInts([]int{1}, 1), sort.SearchInts([]int{1}, 2), sort.SearchInts([]int{1}, 0))

	fmt.Println(searchRange0([]int{1}, 1))

	for i, v := range _34TestCase {
		got := searchRange0(v.in, v.target)
		if !reflect.DeepEqual(got, v.expect) {
			t.Errorf("searchRange [%d].got:%v,wangt:%v", i, got, v.expect)
		}
		got = searchRange1(v.in, v.target)
		if !reflect.DeepEqual(got, v.expect) {
			t.Errorf("searchRange1[%d].got:%v,wangt:%v", i, got, v.expect)
		}
		got = searchRange(v.in, v.target)
		if !reflect.DeepEqual(got, v.expect) {
			t.Errorf("searchRange2[%d].got:%v,wangt:%v", i, got, v.expect)
		}
	}
}
