package util

import (
	"container/list"
	"fmt"
	"strconv"
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

/*
由数组构建二叉树,数组是程序遍历的二叉树
如   	  1_1
          / \
        3_2   2_3
        / \     \
      5_4  3_5   9_7
为[1,3,2,5,3,null,9]    i*2 i*2+1
   0,1,2,3,4, 5,  6
由于有["1", "2", "3", "null", "null", "4", "5", "6", "7"],这种情况,下面的方法要优化,同层次中有null节点且在array中不体现的
无法构建成树
    1
2        3
     4      5
   6   7
TODO: ! 层序数组(leetcode格式的)
*/
func NewFromArrayNotWithLeetcode(strArray []string) (*TreeNode, error) {
	if len(strArray) == 0 {
		return nil, nil
	}
	nodes := make([]*TreeNode, len(strArray))

	for id, str := range strArray {
		node, err := NewNodeFromString(str)
		if err != nil {
			return nil, err
		}
		nodes[id] = node
		if id == 0 {
			continue
		}
		index := id + 1
		if !(index/2 >= 0 && index/2 <= len(nodes)) {
			return nil, fmt.Errorf("index out of range")
		}
		if nodes[(index/2)-1] == nil {
			continue
		}
		if index%2 == 0 {
			nodes[(index/2)-1].Left = node
		} else {
			nodes[(index/2)-1].Right = node
		}
	}
	if len(nodes) > 0 {
		return nodes[0], nil
	}
	return nil, nil
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

// 层序遍历二叉树,空的返回null
// TODO has bug need fix
func (t TreeNode) sourceToFix(root *TreeNode) []string {
	out := []string{}
	if root == nil {
		return out
	}
	ns := []*TreeNode{root}
	for l := len(ns); l > 0; { // 每次l都是1,不知道为什么
		for _, v := range ns {
			// 当前这一层的节点值
			if v == nil {
				out = append(out, NullNode)
				continue
			}
			out = append(out, strconv.Itoa(v.Val))
			// 下一层节点
			if v.Left != nil {
				ns = append(ns, v.Left)
			} else {
				ns = append(ns, nil)
			}
			if v.Right != nil {
				ns = append(ns, v.Right)
			} else {
				ns = append(ns, nil)
			}
		}
		ns = ns[l:]
	}
	return out
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
			nVal := node.Value.(*TreeNode)
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
