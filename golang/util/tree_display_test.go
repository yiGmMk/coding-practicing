package util

import (
	"container/list"
	"fmt"
	"testing"
)

var trees = []struct {
	tree string
}{
	{
		tree: "[0,2,1]",
	},
	{
		tree: "[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]",
	},
	{
		tree: "[0,null,1]",
	},
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
		// fmt.Println(tc.tree, "", root)

		fmt.Println(PrintBinaryTree(root))
	}
}

func PrintBinaryTree(root *TreeNode) string {
	if root == nil {
		return ""
	}
	out := ""
	nodeList := list.New()
	nodeList.PushBack(root)
	for nodeList.Len() > 0 {
		nL := nodeList.Len()
		curLevel := ""
		next := ""
		for i := 0; i < nL; i++ {
			node := nodeList.Front().Value.(*TreeNode)
			nodeList.Remove(nodeList.Front())
			if node == nil {
				continue
			}

			curLevel += fmt.Sprintf("%d", node.Val)
			if node.Left != nil {
				nodeList.PushBack(node.Left)
				next += "/" + " "
			} else {
				next += " "
			}
			if node.Right != nil {
				nodeList.PushBack(node.Right)
				next += "\\" + " "
			} else {
				next += " "
			}
		}
		out += curLevel + "\n" + next + "\n"
	}
	return out
}
