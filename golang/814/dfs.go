package jzoffer

// 深度优先搜索
func fDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left = fDFS(root.Left)
	}
	if root.Right != nil {
		root.Right = fDFS(root.Right)
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
