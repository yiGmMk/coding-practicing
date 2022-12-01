package util

import "strconv"

// Preorder traversal of binary tree
// 前序遍历,根左右
func Tree2PreOrder(root *TreeNode) []string {
	out := []string{}
	if root == nil {
		return out
	}
	out = append(out, strconv.Itoa(root.Val))
	out = append(out, Tree2PreOrder(root.Left)...)
	out = append(out, Tree2PreOrder(root.Right)...)
	return out
}

// Postorder traversal of binary tree
// 后序遍历,左右根
func Tree2PostOrder(root *TreeNode) []string {
	out := []string{}
	if root == nil {
		return out
	}
	out = Tree2PostOrder(root.Left)
	out = append(out, Tree2PostOrder(root.Right)...)
	out = append(out, strconv.Itoa(root.Val))
	return out
}

// Inorder traversal of binary tree
// 中序遍历,左右根
func Tree2InOrder(root *TreeNode) []string {
	out := []string{}
	if root == nil {
		return out
	}
	out = Tree2InOrder(root.Left)
	out = append(out, strconv.Itoa(root.Val))
	out = append(out, Tree2InOrder(root.Right)...)
	return out
}
