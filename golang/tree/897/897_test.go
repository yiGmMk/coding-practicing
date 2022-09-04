package jzoffer

import (
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

func Test897(t *testing.T) {
	for _, tc := range ts {
		root := util.Strs2TreeNode(tc.tree)
		res := increasingBST(root)
		strRes := util.Tree2Strs(res)
		if !reflect.DeepEqual(strRes, tc.res) {
			t.Errorf("increasingBST(%v) = %v, want %v", tc.tree, strRes, tc.res)
		}
	}
}

type testCase struct {
	tree []string
	res  []string
}

var ts = []testCase{
	{
		/*
		 * 输入：root = [5,1,7]
		 * 输出：[1,null,5,null,7]
		 *   5
		 *  / \
		 * 1   7
		 * ------------------------
		 * 1
		 *  \
		 *   5
		 *  / \
		 * null 7
		 */
		tree: []string{"5", "1", "7"},
		res:  []string{"1", "null", "5", "null", "7"},
	},
	{
		/*
		 * 输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
		 * 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
		 *     5
		 *    / \
		 *   3   6
		 *   / \   \
		 *  2   4   8
		 * / \     / \
		 * 1 null  7  9
		 */
		tree: []string{"5", "3", "6", "2", "4", "null", "8", "1", "null", "null", "null", "7", "9"},
		res:  []string{"1", "null", "2", "null", "3", "null", "4", "null", "5", "null", "6", "null", "7", "null", "8", "null", "9"},
	},
}
