/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 *
 * https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (70.14%)
 * Likes:    2508
 * Dislikes: 0
 * Total Accepted:    609.5K
 * Total Submissions: 868K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * 输出：3
 * 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * 输出：5
 * 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], p = 1, q = 2
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [2, 10^5] 内。
 * -10^9
 * 所有 Node.val 互不相同 。
 * p != q
 * p 和 q 均存在于给定的二叉树中。
 *
 *
 */
package jzoffer

import "github.com/yiGmMk/leetcode/golang/util"

type TreeNode = util.TreeNode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// dfs记录每个节点是否是p,q的父节点,bfs找出lowest的那个
func lowestCommonAncestorMy(root, p, q *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (bool, bool)
	resmap := map[*TreeNode][]bool{nil: {false, false}}
	dfs = func(node *TreeNode) (bool, bool) {
		if h, ok := resmap[node]; ok {
			return h[0], h[1]
		}
		if node == p {
			_, has := dfs(node.Left)
			if has {
				resmap[node] = []bool{true, has}
				return true, true
			}
			_, has = dfs(node.Right)
			resmap[node] = []bool{true, has}
			return true, has
		}
		if node == q {
			has, _ := dfs(node.Left)
			if has {
				resmap[node] = []bool{has, true}
				return has, true
			}
			has, _ = dfs(node.Right)
			resmap[node] = []bool{has, true}
			return has, true
		}
		l1, l2 := dfs(node.Left)
		r1, r2 := dfs(node.Right)
		l, r := l1 || r1, l2 || r2
		resmap[node] = []bool{l, r}
		return l, r
	}

	check := func(node *TreeNode) bool {
		if h, ok := resmap[node]; ok && h[0] && h[1] {
			return true
		}
		return false
	}

	dfs(root)

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if check(node) {
			l, r := check(node.Left), check(node.Right)
			if l || r {
				if l {
					queue = append(queue, node.Left)
				}
				if r {
					queue = append(queue, node.Right)
				}
			} else {
				return node
			}
		}
	}
	return root
}

// leetcode
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if r == nil {
		return l
	}
	return r
}

// @lc code=end
