package jzoffer

// 广度有限搜索
// 层序遍历需要从最底层往上删除
// TODO
func fBFS(root *TreeNode) *TreeNode {
	type WithParent struct {
		Node   *TreeNode
		Parent *TreeNode
		left   bool
	}
	if root == nil {
		return root
	}
	queue := []*WithParent{{Node: root, Parent: nil}}
	qL := len(queue)
	for qL > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Node.Left != nil {
				queue = append(queue, &WithParent{Node: node.Node.Left, Parent: node.Node, left: true})
			}
			if node.Node.Right != nil {
				queue = append(queue, &WithParent{Node: node.Node.Right, Parent: node.Node})
			}
			if node.Node.Left == nil && node.Node.Right == nil && node.Node.Val == 0 {
				if node.left {
					node.Parent.Left = nil
				} else {
					node.Parent.Right = nil
				}
			}
		}
		qL = len(queue)
	}
	return root
}
