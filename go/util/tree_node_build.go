package util

import (
	"container/list"
	"strconv"

	"github.com/duke-git/lancet/v2/lancetconstraints"
)

const (
	NullNode = "null" // 空节点
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func (t *TreeNode) AddLeft(node *TreeNode) {
	t.Left = node
}

func (t *TreeNode) AddRight(node *TreeNode) {
	t.Right = node
}

// 字符串数组转二叉树
// references:git@github.com:halfrost/LeetCode-Go.git
func Strs2TreeNode(strs []string) *TreeNode {
	n := len(strs)
	if n == 0 {
		return nil
	}
	root, _ := NewNodeFromString(strs[0])
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && strs[i] != NullNode {
			left, _ := NewNodeFromString(strs[i])
			node.Left = left
			queue = append(queue, left)
		}
		i++
		if i < n && strs[i] != NullNode {
			right, _ := NewNodeFromString(strs[i])
			node.Right = right
			queue = append(queue, right)
		}
		i++
	}
	return root
}

func NewTreeFromString(tree string) *TreeNode {
	nodes := LeetcodeArrayStringToArray(tree)
	root := Strs2TreeNode(nodes)
	return root
}

// 字符串转二叉树节点
func NewNodeFromString(str string) (*TreeNode, error) {
	if str == NullNode {
		return nil, nil
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}
	return &TreeNode{Val: val}, nil
}

/*
由数组构建二叉树,数组是程序遍历的二叉树
如   	  1_1

	    / \
	  3_2   2_3
	  / \     \
	5_4  3_5   9_7

为[1,3,2,5,3,null,9]    i*2 i*2+1

	0,1,2,3,4, 5,  6
*/
func TreeSource(root *TreeNode) []string {
	out := []string{}
	if root == nil {
		return out
	}
	nodeList := list.New()
	nodeList.PushBack(root)
	for nodeList.Len() > 0 {
		nL := nodeList.Len()
		for i := 0; i < nL; i++ {
			node := nodeList.Front()
			nodeList.Remove(node)

			// add到list的都是list的指针不可能为空,需要判断node中的Value
			// if node == nil || node.Value == nil {
			// 	out = append(out, NullNode)
			// 	continue
			// }
			nVal, _ := node.Value.(*TreeNode)
			if nVal == nil {
				out = append(out, NullNode)
				continue
			}
			out = append(out, strconv.Itoa(nVal.Val))

			// 不同时为空,左右子节点必定是要加入的,不管是有值还是nil
			nodeList.PushBack(nVal.Left)
			nodeList.PushBack(nVal.Right)
		}
	}
	// 去除最后一层的NullNode
	i := len(out)
	for i > 0 && out[i-1] == NullNode {
		i--
	}
	return out[:i]
}

// 二叉树,转leetcode格式的层序数组,按行还原
func Tree2Strs(root *TreeNode) []string {
	out := []string{}
	if root == nil {
		return out
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				out = append(out, NullNode)
				continue
			}
			queue = append(queue, node.Left, node.Right)
			out = append(out, strconv.Itoa(node.Val))
		}
		queue = queue[size:]
	}
	i := len(out)
	for i > 0 && out[i-1] == NullNode {
		i--
	}
	return out[:i]
}

// 树的最大深度
// return the max depth of a tree
func (root *TreeNode) Depth() int {
	if root == nil {
		return 0
	}

	return Max(root.Left.Depth(), root.Right.Depth()) + 1
}

// return max value
func Max[T lancetconstraints.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}
