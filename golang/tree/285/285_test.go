package jzoffer

import (
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

var tcs = []struct {
	tree   []string
	target *TreeNode
	p      string
}{
	{
		tree:   []string{"2", "1", "3"},
		p:      "1",
		target: &TreeNode{Val: 2},
	},
	{
		tree:   []string{"5", "3", "6", "2", "4", "null", "null", "1"},
		target: nil,
		p:      "6",
	},
}

func Test285(t *testing.T) {
	for _, tc := range tcs {
		root := util.Strs2TreeNode(tc.tree)
		p := util.Strs2TreeNode([]string{tc.p})
		res := inorderSuccessor3(root, p)
		if res != nil && res.Val != tc.target.Val || (res != nil && tc.target == nil) {
			t.Errorf("got %v, want %v", res, tc.target)
		}
	}
}
