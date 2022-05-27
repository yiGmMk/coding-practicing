package jzoffer

import (
	"container/list"
	"math"
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

// Definition for a binary tree node.

type TreeNode = util.TreeNode

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

var fFixUsingSlice = func(root *TreeNode) []int {
	out := []int{}

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

	f := fFix
	for _, v := range ts {
		tree, err := util.NewFromArray(v.node)
		if err != nil {
			t.Error(err)
			continue
		}
		root := (*TreeNode)(tree)
		res := f(root)
		if !reflect.DeepEqual(res, v.max) {
			t.Errorf("max not equal,expect:%+v,got:%+v", v.max, res)
		}
	}
}
