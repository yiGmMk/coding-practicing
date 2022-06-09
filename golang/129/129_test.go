package jzoffer

import (
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

type TestCase struct {
	tree []string
	res  int
}

var ts = []TestCase{
	{},
	{
		tree: []string{"1", "2", "3"},
		res:  25,
	},
	{
		tree: []string{"4", "9", "0", "5", "1"},
		res:  1026, // 495+491+40
	},
}

func Test129(t *testing.T) {
	for _, tc := range ts {
		tree := util.Strs2TreeNode(tc.tree)
		if tree == nil && len(tc.tree) != 0 {
			continue
		}

		res := sumNumbers(tree)
		if res != tc.res {
			t.Errorf("sumNumbers(%v) = %v, want %v", tc.tree, res, tc.res)
		}
	}
}
