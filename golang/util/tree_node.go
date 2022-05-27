package util

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
