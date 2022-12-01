package jzoffer

import (
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

func pathSum01(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
	}
	dfs(root, 0)
	return
}

type TestCase struct {
	tree   []string
	res    int
	target int
}

func Test437(t *testing.T) {
	ts := []TestCase{
		{
			/*     10
			 *     /  \
			 *    5   -3
			 *   / \   \
			 *  3   2   11
			 * / \   \
			 * 3  -2   1
			 */
			tree:   []string{"10", "5", "-3", "3", "2", "null", "11", "3", "-2", "null", "1"},
			target: 8,
			res:    3,
		},
	}

	fFix := pathSum
	for _, tc := range ts {
		root := util.Strs2TreeNode(tc.tree)
		res := fFix(root, tc.target)
		if res != tc.res {
			t.Errorf("expected: %v, got: %v", tc.res, res)
		}
	}
}
