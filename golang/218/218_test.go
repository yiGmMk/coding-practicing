package jzoffer

import (
	"reflect"
	"testing"
)

var _218TestCase = []struct {
	input  [][]int
	expect [][]int
}{
	{
		input: [][]int{
			{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8},
		},
		expect: [][]int{
			{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0},
		},
	},
}

func Test218(t *testing.T) {
	for i, v := range _218TestCase {
		res := getSkylineOutOfMemory(v.input)
		if !reflect.DeepEqual(res, v.expect) {
			t.Errorf("[%d]not equal,want:%+v,got:%+v", i, v.expect, res)
		}
	}
}
