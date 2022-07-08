package util

import (
	"fmt"
	"strings"
	"testing"
)

var trees = []struct {
	tree string
}{
	{
		tree: "[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]",
	},
	{
		tree: "[0,2,1]",
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

		fmt.Println(printTree1(root))

		fmt.Println(root.PrintHorizontally())
		//d := root.Depth()
		// fmt.Println(tc.tree, "", root, "tree-depth:", d, "")

		//printTree(strs, d)
		// fmt.Println(PrintBinaryTree(root))

	}
}

// TODO, print tree like
/*  0
 * / \
 * 1   2
 */
func printTree(ns []string, depth int) {
	out := strings.Builder{}
	i, j, depth := 0, 0, depth
	for j = 0; j < depth; j++ {
		w := 1 << (depth - j + 1)
		for i = 0; i < 1<<j; i++ {
			out.WriteString(strings.Repeat(" ", w) + "0" + strings.Repeat(" ", w) + " ")
		}
		out.WriteString("\n")
	}
	fmt.Println(out.String())
}

// func PrintBinaryTree(root *TreeNode) string {
// 	if root == nil {
// 		return ""
// 	}
// 	out := ""
// 	nodeList := list.New()
// 	nodeList.PushBack(root)
// 	levels := []*TreeNode{}
// 	for nodeList.Len() > 0 {
// 		nL := nodeList.Len()
// 		for i := 0; i < nL; i++ {
// 			node := nodeList.Front().Value.(*TreeNode)
// 			nodeList.Remove(nodeList.Front())
// 			if node == nil {
// 				continue
// 			}
// 			if node.Left != nil {
// 				nodeList.PushBack(node.Left)
// 				next += "/" + " "
// 			}
// 			if node.Right != nil {
// 				nodeList.PushBack(node.Right)
// 				next += "\\" + " "
// 			}
// 		}
// 		out += curLevel + "\n" + next + "\n"
// 	}
// 	return out
// }

/*
可以分配a[2^d-1]={NULL} 数组，用dfs(node, m) 遍历，
从 dfs(root, 1) 开始，
填入a[m] = node，
递归 dfs(node->left, m * 2);
dfs(node->right, m * 2 + 1)。
那么画顺序第 m 个节点及其向父的分支时，若 a[m] 非空则画出节点和分支，否则画空格。
 0
1 2
[0,"","",1,] len=2^3-1=7
int main() {
	int i, j, k, depth = 4;
	for (j = 0; j < depth; j++) {
		int w = 1 << (depth - j + 1);
		if (j == 0)
			printf("%*c\n", w, '_');
		else {
			for (i = 0; i < 1 << (j - 1); i++) {
				printf("%*c", w + 1, ' ');
				for (k = 0; k < w - 3; k++) printf("_");
				printf("/ \");
				for (k = 0; k < w - 3; k++) printf("_");
				printf("%*c", w + 2, ' ');
			}
			printf("\n");
			for (i = 0; i < 1 << (j - 1); i++)
				printf("%*c/%*c_%*c", w, '_', w * 2 - 2, '\', w, ' ');
			printf("\n");
		}
		for (i = 0; i < 1 << j; i++)
			printf("%*c_)%*c", w - 1, '(', w - 1, ' ');
		printf("\n");
	}
}
*/
//| | | | | | |1| | | | | |
//| | | | | |/| |\| | | | |
//| | | | |/| | | |\| | | |
//| | | |0| | | | | |2| | |
//| | |/| |\| | | |/| |\| |
//| |1| | | |1| |0| | | |1|

// #include <stdio.h>

// int main() {
// 	int i, j, depth = 4;
// 	for (j = 0; j < depth; j++) {
// 		int w = 1 << (depth - j + 1);
// 		for (i = 0; i < 1 << j; i++)
// 			printf("%*c%*c", w, 'o', w, ' ');
// 		printf("\n");
// 	}
// }
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
