package jzoffer

import (
	"reflect"
	"testing"
)

var _658TestCase = []struct {
	arr    []int
	k, x   int
	expect []int
}{
	{
		arr:    []int{1, 2, 3, 4, 5},
		k:      4,
		x:      3,
		expect: []int{1, 2, 3, 4},
	},
	{
		arr:    []int{1, 2, 3, 4, 5},
		k:      4,
		x:      -1,
		expect: []int{1, 2, 3, 4},
	},
}

func Test658(t *testing.T) {
	for i, v := range _658TestCase {
		got := findClosestElements(v.arr, v.k, v.x)
		if !reflect.DeepEqual(v.expect, got) {
			t.Errorf("[%d].got:%v,want:%v", i, got, v.expect)
		}

		got = findClosestElements1(v.arr, v.k, v.x)
		if !reflect.DeepEqual(v.expect, got) {
			t.Errorf("[%d]1.got:%v,want:%v", i, got, v.expect)
		}
	}
}
