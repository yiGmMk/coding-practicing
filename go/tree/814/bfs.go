package jzoffer

// 广度有限搜索
// 层序遍历需要从最底层往上删除
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
	levels := [][]*WithParent{}
	qL := len(queue)
	for qL > 0 {
		n := len(queue)
		level := []*WithParent{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Node.Left != nil {
				queue = append(queue, &WithParent{Node: node.Node.Left, Parent: node.Node, left: true})
			}
			if node.Node.Right != nil {
				queue = append(queue, &WithParent{Node: node.Node.Right, Parent: node.Node})
			}
			level = append(level, node)
		}
		qL = len(queue)
		levels = append(levels, level)
	}
	for i := len(levels) - 1; i >= 0; i-- {
		for _, n := range levels[i] {
			if n.Node.Left == nil && n.Node.Right == nil && n.Node.Val == 0 {
				if n.Parent == nil {
					return nil // 没有parent的是root,root也符合条件直接返回nil
				}
				if n.left {
					n.Parent.Left = nil
				} else {
					n.Parent.Right = nil
				}
			}
		}
	}
	return root
}
