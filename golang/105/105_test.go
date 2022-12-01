package jzoffer

import (
	"log"
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

type testCase struct {
	preorder []string
	inOrder  []string
	res      []string
}

var ts = []testCase{
	// * 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// * 输出: [3,9,20,null,null,15,7]
	{
		preorder: []string{"3", "9", "20", "15", "7"},
		inOrder:  []string{"9", "3", "15", "20", "7"},
		res:      []string{"3", "9", "20", "null", "null", "15", "7"},
	},
}

var fFix func(preorder, inorder []int) *TreeNode

func Test105(t *testing.T) {
	fFix = buildTree

	for _, tc := range ts {
		pre, _ := util.Str2Int(tc.preorder)
		in, _ := util.Str2Int(tc.inOrder)
		root := fFix(pre, in)
		strRes := util.Tree2Strs(root)
		if reflect.DeepEqual(strRes, tc.res) {
			log.Printf("res not equal,expect:%+v,got:%+v\n", tc.res, root)
		}
	}
}
