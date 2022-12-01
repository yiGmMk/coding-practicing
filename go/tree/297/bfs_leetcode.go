package jzoffer

import (
	"strconv"
	"strings"
)

type BfsLeetcode struct {
}

const (
	NullNode = "null" // 空节点
)

func (b *BfsLeetcode) Serialize(root *TreeNode) string {
	out := []string{}
	if root == nil {
		return ""
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
	return strings.Join(out[:i], ",")
}

func (b *BfsLeetcode) Deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
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
