package util

import "fmt"

func (node *TreeNode) String() string {
	str := "Binary Tree:\n"
	if node != nil {
		treeOutput(node, "", true, &str)
	}
	return str
}

func (node *TreeNode) val() string {
	return fmt.Sprintf("%v", node.Val)
}

// 输出二叉树,打印到控制台
func treeOutput(node *TreeNode, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		treeOutput(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.val() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		treeOutput(node.Left, newPrefix, true, str)
	}
}
