/*
 * @lc app=leetcode.cn id=1372 lang=golang
 *
 * [1372] 二叉树中的最长交错路径
 *
 * https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/description/
 *
 * algorithms
 * Medium (55.45%)
 * Likes:    153
 * Dislikes: 0
 * Total Accepted:    22K
 * Total Submissions: 39.6K
 * Testcase Example:  '[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]'
 *
 * 给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：
 *
 *
 * 选择二叉树中 任意 节点和一个方向（左或者右）。
 * 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
 * 改变前进方向：左变右或者右变左。
 * 重复第二步和第三步，直到你在树中无法继续移动。
 *
 *
 * 交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
 *
 * 请你返回给定树中最长 交错路径 的长度。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
 * 输出：3
 * 解释：蓝色节点为树中最长交错路径（右 -> 左 -> 右）。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：root = [1,1,1,null,1,null,null,1,1,null,1]
 * 输出：4
 * 解释：蓝色节点为树中最长交错路径（左 -> 右 -> 左 -> 右）。
 *
 *
 * 示例 3：
 *
 * 输入：root = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每棵树最多有 50000 个节点。
 * 每个节点的值在 [1, 100] 之间。
 *
 *
 */
package jzoffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO
func longestZigZagDfs(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}

	var (
		// left 指定下一个是什么节点时+1
		dfs func(root *TreeNode, left bool, length int)
	)
	res = 1 // 路径上的节点数
	dfs = func(root *TreeNode, left bool, length int) {
		if root == nil {
			return
		}
		if left {
			if root.Left != nil {
				res = max(res, length+1)
				dfs(root.Left, false, length+1)
			}
			// 下一个必须是右节点才能加1
			if root.Right != nil {
				dfs(root.Right, false, 1)
			}
			return
		}
		if root.Right != nil {
			res = max(res, length+1)
			dfs(root.Right, true, length+1)
		}
		if root.Left != nil {
			dfs(root.Left, true, 1)
		}
	}
	dfs(root, true, 1)
	dfs(root, false, 1)
	return res - 1
}

// dp 动态规划
func longestZigZag(root *TreeNode) (res int) {
	var (
		f, g  = map[*TreeNode]int{root: 0}, map[*TreeNode]int{root: 0}
		queue = [][]*TreeNode{{root, nil}}
	)
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		father, child := pair[1], pair[0]
		f[child], g[child] = 0, 0
		if father != nil {
			if father.Left == child {
				f[child] = g[father] + 1
			}
			if father.Right == child {
				g[child] = f[father] + 1
			}
		}
		if child.Left != nil {
			queue = append(queue, []*TreeNode{child.Left, child})
		}
		if child.Right != nil {
			queue = append(queue, []*TreeNode{child.Right, child})
		}
	}
	for k, v := range f {
		res = max(res, max(v, g[k]))
	}
	return
}

// @lc code=end
