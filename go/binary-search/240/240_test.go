package jzoffer

import "testing"

var _240TestCase = []struct {
	vs     [][]int
	target int
	expect bool
}{
	// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
	// 输出：true
	// 示例 2：
	// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
	// 输出：false
	{
		vs:     [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		target: 3,
		expect: true,
	},
	{
		vs:     [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		target: 13,
		expect: false,
	},
	{
		vs:     [][]int{{1}},
		target: 0,
		expect: false,
	},
	{
		vs:     [][]int{{1}},
		target: 1,
		expect: true,
	},
}

func Test240(t *testing.T) {
	for i, v := range _240TestCase {
		got := searchMatrix(v.vs, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%v,want:%v", i, got, v.expect)
		}
		got = searchMatrix1(v.vs, v.target)
		if got != v.expect {
			t.Errorf("[%d]1.got:%v,want:%v", i, got, v.expect)
		}
	}
}
