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
*/
func NewFromArray(strArray []string) (*TreeNode, error) {
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

			// 叶子节点不要加null了,不然多一层null,无意义
			if nVal.Left == nil && nVal.Right == nil {
				continue
			}
			// 不同时为空,左右子节点必定是要加入的,不管是有值还是nil
			nodeList.PushBack(nVal.Left)
			nodeList.PushBack(nVal.Right)
		}
	}

	return out
}
