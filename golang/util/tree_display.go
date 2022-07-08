package util

import (
	"fmt"
	"strings"
)

/* many ways to print a tree
 * here print tree [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8] like this:
 * \---4
 *     |---1
 *     |   |---0
 *     |   \---2
 *     |       \---3
 *     \---6
 *         |---5
 *         \---7
 *             \---8
 */
func (root *TreeNode) String() string {
	return printTree1(root)
}

/* display tree like this:
 *  BinaryTree
 * │           ┌── 6
 * │       ┌── 5
 * │   ┌── 4
 * │   │   └── 3
 * └── 2
 *     └── 1
 */
func (root *TreeNode) PrintHorizontally() string {
	output := "BinaryTree\n"
	printTree2(root, "", true, &output)
	return output
}

/* references:https://stackoverflow.com/questions/4965335/how-to-print-binary-tree-diagram-in-java
 * print tree like this:
 *   0
 *  / \
 * 1   2
 */
func (root *TreeNode) PrintVertically() string {
	return printTree3(root)
}

func (node *TreeNode) val() string {
	return fmt.Sprintf("%v", node.Val)
}

/* print tree like [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
 * \---4
 *     |---1
 *     |   |---0
 *     |   \---2
 *     |       \---3
 *     \---6
 *         |---5
 *         \---7
 *             \---8
 * swift:
 * func printTree(prefix: String, n: TreeNode, isLeft: Bool) {
 *        print(prefix, (isLeft ? "|-- " : "\\-- "), n.value)
 *        if n.left != nil {
 *            printTree(prefix: "\(prefix) \(isLeft ? "|   " : "   ") ", n: n.left!, isLeft: true)
 *        }
 *        if n.right != nil {
 *            printTree(prefix: "\(prefix) \(isLeft ? "|   " : "   ") ", n: n.right!, isLeft: false)
 *        }
 *    }
 * references: https://stackoverflow.com/questions/43898440/how-to-draw-a-binary-tree-in-console
 */
func printTree1(root *TreeNode) string {
	builder := strings.Builder{}
	var print func(root *TreeNode, prefix string, isLeft bool)

	print = func(root *TreeNode, prefix string, isLeft bool) {
		if isLeft {
			builder.WriteString(fmt.Sprintf("%s|---%d\n", prefix, root.Val))
		} else {
			builder.WriteString(fmt.Sprintf("%s\\---%d\n", prefix, root.Val))
		}
		padding := "    "
		if isLeft {
			padding = "|   "
		}
		if root.Left != nil {
			print(root.Left, prefix+padding, true)
		}
		if root.Right != nil {
			print(root.Right, prefix+padding, false)
		}
	}
	print(root, "", false)

	return builder.String()
}

/*
 * print tree like this,copy code from:github.com/emirpasic/gods/trees/redblacktree
 * 输出二叉树,打印到控制台
 *  BinaryTree
 * │           ┌── 6
 * │       ┌── 5
 * │   ┌── 4
 * │   │   └── 3
 * └── 2
 *     └── 1
 */
func printTree2(node *TreeNode, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		printTree2(node.Right, newPrefix, false, str)
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
		printTree2(node.Left, newPrefix, true, str)
	}
}

func Repeat(in string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(in, count)
}

func printTree3(root *TreeNode) string {
	builder := strings.Builder{}
	var print func(root []*TreeNode, cur, maxDepth int)
	print = func(nodes []*TreeNode, cur, maxDepth int) {
		if len(nodes) == 0 || isAllNull(nodes) || cur > maxDepth {
			return
		}
		floor := maxDepth - cur
		endgeLines := 1 << (Max(floor-1, 0))
		firstSpaces := 1<<(floor) - 1
		betweenSpaces := 1<<(floor+1) - 1 //

		builder.WriteString(Repeat(" ", firstSpaces))

		// print node values
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node != nil {
				builder.WriteString(fmt.Sprintf("%d", node.Val))
				newNodes = append(newNodes, node.Left)
				newNodes = append(newNodes, node.Right)
			} else {
				newNodes = append(newNodes, nil, nil)
				builder.WriteString(" ")
			}

			builder.WriteString(Repeat(" ", betweenSpaces))
		}
		builder.WriteString("\n")

		// edges between levels
		for i := 1; i <= endgeLines; i++ {
			for j := 0; j < len(nodes); j++ {
				builder.WriteString(Repeat(" ", firstSpaces-i))
				if nodes[j] == nil {
					builder.WriteString(Repeat(" ", endgeLines+endgeLines+i+1))
					continue
				}

				if nodes[j].Left != nil {
					builder.WriteString("/")
				} else {
					builder.WriteString(" ")
				}

				builder.WriteString(Repeat(" ", i+i-1))

				if nodes[j].Right != nil {
					builder.WriteString("\\")
				} else {
					builder.WriteString(" ")
				}
				builder.WriteString(Repeat(" ", endgeLines+endgeLines-i))
			}
			builder.WriteString("\n")
		}
		print(newNodes, cur+1, maxDepth)
	}
	print([]*TreeNode{root}, 1, root.Depth())
	return builder.String()
}

func isAllNull(nodes []*TreeNode) bool {
	for _, n := range nodes {
		if n != nil {
			return false
		}
	}
	return true
}
