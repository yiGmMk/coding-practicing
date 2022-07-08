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

func (root *TreeNode) PrintHorizontally() string {
	output := "BinaryTree\n"
	printTree2(root, "", true, &output)
	return output
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
 *  RedBlackTree
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

/*
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class BTreePrinterTest {
class Node<T extends Comparable<?>> {
    Node<T> left, right;
    T data;

    public Node(T data) {
        this.data = data;
    }
}

class BTreePrinter {
    public static <T extends Comparable<?>> void printNode(Node<T> root) {
        int maxLevel = BTreePrinter.maxLevel(root);
        printNodeInternal(Collections.singletonList(root), 1, maxLevel);
    }

    private static <T extends Comparable<?>> void printNodeInternal(List<Node<T>> nodes, int level, int maxLevel) {
        if (nodes.isEmpty() || BTreePrinter.isAllElementsNull(nodes))
            return;

        int floor = maxLevel - level;
        int endgeLines = (int) Math.pow(2, (Math.max(floor - 1, 0)));
        int firstSpaces = (int) Math.pow(2, (floor)) - 1;
        int betweenSpaces = (int) Math.pow(2, (floor + 1)) - 1;

        BTreePrinter.printWhitespaces(firstSpaces);

        List<Node<T>> newNodes = new ArrayList<Node<T>>();
        for (Node<T> node : nodes) {
            if (node != null) {
                System.out.print(node.data);
                newNodes.add(node.left);
                newNodes.add(node.right);
            } else {
                newNodes.add(null);
                newNodes.add(null);
                System.out.print(" ");
            }

            BTreePrinter.printWhitespaces(betweenSpaces);
        }
        System.out.println("");

        for (int i = 1; i <= endgeLines; i++) {
            for (int j = 0; j < nodes.size(); j++) {
                BTreePrinter.printWhitespaces(firstSpaces - i);
                if (nodes.get(j) == null) {
                    BTreePrinter.printWhitespaces(endgeLines + endgeLines + i + 1);
                    continue;
                }

*/
