package jzoffer

import (
	"math"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

type testCase struct {
	tree []string
	res  int
}

func Test124(t *testing.T) {
	ts := []testCase{
		{
			tree: []string{"-10", "9", "20", "null", "null", "15", "7"},
			res:  42,
		},
	}
	for _, tc := range ts {
		root := util.Strs2TreeNode(tc.tree)
		res := maxPathSum01(root)
		if res != tc.res {
			t.Errorf("expected: %v, got: %v", tc.res, res)
		}
	}
}

func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxL := max(maxGain(root.Left), 0)
	maxR := max(maxGain(root.Right), 0)

	priceNewPath := maxL + maxR + root.Val
	maxSum = max(maxSum, priceNewPath)
	return max(maxL, maxR) + root.Val
}

var maxSum = math.MinInt

func maxPathSum01(root *TreeNode) int {
	// 多测测试会影响全局变量,在leetcode提交时需要初始化
	maxSum = math.MinInt
	maxGain(root)
	return maxSum
}
