package jzoffer

import (
	"fmt"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

func Test236(t *testing.T) {
	root := util.NewTreeFromString("[3,5,1,6,2,0,8,null,null,7,4]")

	fmt.Println(lowestCommonAncestorMy(root, root.Left, root.Right))
}
