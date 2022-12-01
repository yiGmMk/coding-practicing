package jzoffer

import (
	"strings"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

var _226TestCase = []struct {
	tree, res string
}{
	/*
	 *输入：root = [4,2,7,1,3,6,9]
	 *输出：[4,7,2,9,6,3,1]
	 */
	{
		tree: "[4,2,7,1,3,6,9]",
		res:  "[4,7,2,9,6,3,1]",
	},
}

func Test226(t *testing.T) {
	for _, v := range _226TestCase {
		root := util.NewTreeFromString(v.tree)
		res := invertTree(root)
		if got := "[" + strings.Join(util.Tree2Strs(res), ",") + "]"; got != v.res {
			t.Errorf("got:%s,expect:%s", got, v.res)
		}
	}
}
