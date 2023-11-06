package tree

import "strings"

type node struct {
	val  string
	l, r *node
}

func postTree(pre, mid string) string {
	var dfs func(p, m string) *node
	dfs = func(p, m string) *node {
		if p == "" {
			return nil
		}
		if len(p) == 1 {
			return &node{val: p}
		}
		bRoot := p[0]
		// 从中序遍历找根节点位置
		i := strings.IndexByte(m, bRoot) // [0:i][i+1:]
		pl, ml := p[1:i+1], m[:i]
		pr, mr := p[i+1:], m[i+1:]
		root := &node{
			val: string(bRoot),
			l:   dfs(pl, ml), // 左子树
			r:   dfs(pr, mr),
		}

		return root
	}
	root := dfs(pre, mid)
	if root == nil {
		return ""
	}

	b := strings.Builder{}
	var post func(root *node)
	post = func(root *node) {
		if root == nil {
			return
		}
		if root.l != nil {
			post(root.l)
		}
		if root.r != nil {
			post(root.r)
		}
		b.WriteString(root.val)
	}
	post(root)
	return b.String()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 基于前/中序遍历 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(p, m []int) *TreeNode
	dfs = func(p, m []int) *TreeNode {
		if p == nil || len(p) == 0 {
			return nil
		}
		if len(p) == 1 {
			return &TreeNode{Val: p[0]}
		}
		bRoot := p[0]
		i := 0
		for j, v := range m {
			if v == bRoot {
				i = j
				break
			}
		}
		pl, ml := p[1:i+1], m[:i]
		pr, mr := p[i+1:], m[i+1:]
		root := &TreeNode{
			Val:   bRoot,
			Left:  dfs(pl, ml), // 左子树
			Right: dfs(pr, mr),
		}

		return root
	}
	root := dfs(preorder, inorder)
	return root
}

// 后续遍历
func postItrate(root *TreeNode) (res []int) {
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
	dfs(root)
	return
}
