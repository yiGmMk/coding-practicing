package jzoffer

import (
	"container/list"
	"math"
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

// 从上往下,层序遍历数,每层计算最大值,加入数据
var fFix = func(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	level := list.New()
	level.PushBack(root)
	for level.Len() > 0 {
		nl := level.Len()
		max := math.MinInt32
		for i := 0; i < nl; i++ {
			node := level.Front()
			level.Remove(node)
			val := node.Value.(*TreeNode)
			if val == nil {
				continue
			}
			max = Max(max, val.Val)
			left := val.Left
			if left != nil {
				level.PushBack(left)
			}
			right := val.Right
			if right != nil {
				level.PushBack(right)
			}
		}
		out = append(out, max)
	}
	return out
}

// 使用切片实现
var fFixUsingSlice = func(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	level := []*TreeNode{root}
	levelNum := len(level)
	for levelNum > 0 {
		max := math.MinInt32
		for i := 0; i < levelNum; i++ {
			node := level[0]
			level = level[1:]
			if node == nil {
				continue
			}
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		out = append(out, max)
		levelNum = len(level)
	}
	return out
}

// test 515 solution
func Test515(t *testing.T) {
	ts := []struct {
		node []string
		max  []int
	}{
		{[]string{"1", "3", "2", "5", "3", "null", "9"}, []int{1, 3, 9}},
	}

	f := fFixUsingSlice
	for _, v := range ts {
		tree := util.Strs2TreeNode(v.node)
		root := (*TreeNode)(tree)
		res := f(root)
		if !reflect.DeepEqual(res, v.max) {
			t.Errorf("max not equal,expect:%+v,got:%+v", v.max, res)
		}
	}
}
