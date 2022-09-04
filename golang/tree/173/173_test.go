package jzoffer

import (
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

func Test173(t *testing.T) {
	treeStr := "[7, 3, 15, null, null, 9, 20]"
	strNodes := util.LeetcodeArrayStringToArray(treeStr)
	root := util.Strs2TreeNode(strNodes)

	/*
		inputs = ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
		inputs = [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
		输出
		[null, 3, 7, true, 9, true, 15, true, 20, false]
	*/
	bst := Constructor(root)
	if bst.Next() != 3 {
		t.Errorf("next error")
	}
	if bst.Next() != 7 {
		t.Errorf("next error")
	}
	if bst.HasNext() != true {
		t.Errorf("hasNext error")
	}
	if bst.Next() != 9 {
		t.Errorf("next error")
	}
	if bst.HasNext() != true {
		t.Errorf("hasNext error")
	}
	if bst.Next() != 15 {
		t.Errorf("next error")
	}
	if bst.HasNext() != true {
		t.Errorf("hasNext error")
	}
	if bst.Next() != 20 {
		t.Errorf("next error")
	}
	if bst.HasNext() != false {
		t.Errorf("hasNext error")
	}
}
