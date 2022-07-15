package util

import (
	"fmt"
	"testing"
)

var trees = []struct {
	tree string
}{
	{
		tree: "[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]",
	},
	// {
	// 	tree: "[0,2,1]",
	// },

	// {
	// 	tree: "[0,null,1]",
	// },
}

/*
    3
   / \
  9  20
    /  \
   15   7

"Binary Tree:\n│   ┌── 1\n└── 0\n    └── 2\n"
*/
func TestDisplayTree(t *testing.T) {

	for _, tc := range trees {
		strs := LeetcodeArrayStringToArray(tc.tree)
		root := Strs2TreeNode(strs)

		fmt.Println(root)

		fmt.Println(root.PrintHorizontally())

		fmt.Println(root.PrintVertically())
	}
}

// TODO, print tree like
/*  0
 * / \
 * 1   2
 */
// func printTree(ns []string, depth int) {
// 	out := strings.Builder{}
// 	i, j, depth := 0, 0, depth
// 	for j = 0; j < depth; j++ {
// 		w := 1 << (depth - j + 1)
// 		for i = 0; i < 1<<j; i++ {
// 			out.WriteString(strings.Repeat(" ", w) + "0" + strings.Repeat(" ", w) + " ")
// 		}
// 		out.WriteString("\n")
// 	}
// 	fmt.Println(out.String())
// }
