package jzoffer

import (
	"testing"
)

var _19TestCase = []struct {
	list string
	n    int
	res  string
}{
	{
		list: "[1,2,3,4,5]",
		n:    2,
		res:  "[1,2,3,5]",
	},

	/*
	 * 输入：head = [1], n = 1
	 * 输出：[]
	 *
	 *
	 * 示例 3：
	 *
	 *
	 * 输入：head = [1,2], n = 1
	 * 输出：[1]*/
	{
		list: "[1]",
		n:    1,
		res:  "[]",
	},
	{
		list: "[1,2]",
		n:    1,
		res:  "[1]",
	},
}

func Test19(t *testing.T) {

}
