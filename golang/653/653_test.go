package jzoffer

import (
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

var ts = []struct {
	tree   string
	target int
	res    bool
}{
	{
		tree:   "[2,1,3,-4,null,null,null,null,0]",
		target: 2,
		res:    true,
	},
	{
		tree:   "[2,0,3,-4,1]",
		target: -1,
		res:    true,
	},
	{
		tree:   "[1]",
		target: 2,
		res:    false,
	},
	{
		tree:   "[5,3,6,2,4,null,7]",
		target: 9,
		res:    true,
	},
	{
		/*
		* 输入: root = [5,3,6,2,4,null,7], k = 28
		* 输出: false
		 */
		tree:   "[5,3,6,2,4,null,7]",
		target: 28,
		res:    false,
	},
}

func Test653(t *testing.T) {
	for _, tc := range ts {
		strArr := util.LeetcodeArrayStringToArray(tc.tree)
		root := util.Strs2TreeNode(strArr)
		if findTarget(root, tc.target) != tc.res {
			t.Errorf("findTarget != %v", tc.res)
		}
	}
}
