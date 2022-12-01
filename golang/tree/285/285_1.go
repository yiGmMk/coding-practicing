package jzoffer

// 方法一：中序遍历
// 为了找到二叉搜索树中的节点 pp 的中序后继，最直观的方法是中序遍历。由于只需要找到节点 pp 的中序后继，
// 因此不需要维护完整的中序遍历序列，只需要在中序遍历的过程中维护上一个访问的节点和当前访问的节点。
// 如果上一个访问的节点是节点 pp，则当前访问的节点即为节点 pp 的中序后继。

// 如果节点 pp 是最后被访问的节点，则不存在节点 pp 的中序后继，返回 \text{null}null。
// 复杂度分析
// 时间复杂度：O(n)，其中 n 是二叉搜索树的节点数。中序遍历最多需要访问二叉搜索树中的每个节点一次。
// 空间复杂度：O(n)，其中 n 是二叉搜索树的节点数。空间复杂度取决于栈深度，平均情况是 O(\log n)，最坏情况是 O(n)。
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	st := []*TreeNode{}
	var pre, cur *TreeNode = nil, root
	for len(st) > 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	// 2
	//1  3 =>
	return nil
}

// 方法二：利用二叉搜索树的性质
// 二叉搜索树的一个性质是中序遍历序列单调递增，因此二叉搜索树中的节点 pp 的中序后继满足以下条件：
// 中序后继的节点值大于 pp 的节点值；
// 中序后继是节点值大于 pp 的节点值的所有节点中节点值最小的一个节点。
// 利用二叉搜索树的性质，可以在不做中序遍历的情况下找到节点 pp 的中序后继。
// 如果节点 pp 的右子树不为空，则节点 pp 的中序后继在其右子树中，在其右子树中定位到最左边的节点，即为节点 pp 的中序后继。
// 如果节点 pp 的右子树为空，则需要从根节点开始遍历寻找节点 pp 的祖先节点。
// 将答案初始化为 \text{null}null。用 \textit{node}node 表示遍历到的节点，初始时 \textit{node} = \textit{root}node=root。每次比较 \textit{node}node 的节点值和 pp 的节点值，执行相应操作：

// 如果 \textit{node}node 的节点值大于 pp 的节点值，则 pp 的中序后继可能是 \textit{node}node 或者在 \textit{node}node 的左子树中，因此用 \textit{node}node 更新答案，并将 \textit{node}node 移动到其左子节点继续遍历；

// 如果 \textit{node}node 的节点值小于或等于 pp 的节点值，则 pp 的中序后继可能在 \textit{node}node 的右子树中，因此将 \textit{node}node 移动到其右子节点继续遍历。

// 由于在遍历过程中，当且仅当 \textit{node}node 的节点值大于 pp 的节点值的情况下，才会用 \textit{node}node 更新答案，因此当节点 pp 有中序后继时一定可以找到中序后继，当节点 pp 没有中序后继时答案一定为 \text{null}null。

func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}

// 复杂度分析

// 时间复杂度：O(n)O(n)，其中 nn 是二叉搜索树的节点数。遍历的节点数不超过二叉搜索树的高度，平均情况是 O(\log n)O(logn)，最坏情况是 O(n)O(n)。

// 空间复杂度：O(1)O(1)。

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/P5rCT8/solution/er-cha-sou-suo-shu-zhong-de-zhong-xu-hou-5nrz/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
