package jzoffer

import (
	"container/list"
	"testing"
)

// Definition for a binary tree node.

// 从上往下,层序遍历数,每层计算最大值,加入数据
var fFix = func(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	level := list.New()
	level.PushBack(root)
	for n := level.Len(); n > 0; {
		max := level.Front().Value.(*TreeNode).Val
		for i := 0; i < n; i++ {
			val := level.Front()
			if val == nil {
				continue
			}
			max = Max(max, val.Value.(*TreeNode).Val)
			left := val.Value.(*TreeNode).Left
			if left != nil {
				level.PushBack(left)
			}
			right := val.Value.(*TreeNode).Right
			if right != nil {
				level.PushBack(right)
			}
			level.Remove(val)
		}
		out = append(out, max)
	}
	return out
}

func Test515(t *testing.T) {

}
