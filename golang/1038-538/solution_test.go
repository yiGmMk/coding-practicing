package jzoffer

import (
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

var ts = []struct {
	tree []string
	res  []string
}{
	{
		tree: []string{"2", "1", "3"},
		res:  []string{"5", "6", "3"},
	},
	// 输入：root = [0,null,1]
	// 输出：[1,null,1]
	{
		tree: []string{"0", "null", "1"},
		res:  []string{"1", "null", "1"},
	},

	// 输入：root = [1,0,2]
	// 输出：[3,3,2]
	{
		tree: []string{"1", "0", "2"},
		res:  []string{"3", "3", "2"},
	},

	// 输入：root = [3,2,4,1]
	// 输出：[7,9,4,10]
	{
		tree: []string{"3", "2", "4", "1"},
		res:  []string{"7", "9", "4", "10"},
	},

	{
		/*
			输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
			输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
		*/
		tree: []string{"4", "1", "6", "0", "2", "5", "7", "null", "null", "null", "3", "null", "null", "null", "8"},
		res:  []string{"30", "36", "21", "36", "35", "26", "15", "null", "null", "null", "33", "null", "null", "null", "8"},
	},
}

func Test1038(t *testing.T) {
	for _, tc := range ts {
		root := util.Strs2TreeNode(tc.tree)
		restree := convertBST(root)
		if len(tc.res) == 0 && restree == nil {
			continue
		}
		arrayRes := util.Tree2Strs(restree)
		if !reflect.DeepEqual(arrayRes, tc.res) {
			t.Errorf("got %v, want %v", arrayRes, tc.res)
		}
	}
}
