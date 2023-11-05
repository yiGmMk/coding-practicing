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
		i := strings.IndexByte(p, bRoot) // [0:i][i+1:]
		root := &node{
			val: string(bRoot),
			l:   dfs(p[1:i+1], m[:i]), // 左子树
			r:   dfs(p[i+1:], m[i+1:]),
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
		root := &TreeNode{
			Val:   bRoot,
			Left:  dfs(p[1:i+1], m[:i]), // 左子树
			Right: dfs(p[i+1:], m[i+1:]),
		}

		return root
	}
	root := dfs(preorder, inorder)
	return root
}
