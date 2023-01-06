package jzoffer

import (
	"reflect"
	"testing"
)

var _704TestCase = []struct {
	nums   []int
	target int
	expect int
}{
	{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 9,
		expect: 4,
	},
	{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 2,
		expect: -1,
	},
	{
		nums:   []int{5},
		target: -5,
		expect: -1,
	},
}

func Test704(t *testing.T) {
	for i, v := range _704TestCase {
		got := search(v.nums, v.target)
		if !reflect.DeepEqual(v.expect, got) {
			t.Errorf("[%d].got:%v,want:%v", i, got, v.expect)
		}
	}
}
