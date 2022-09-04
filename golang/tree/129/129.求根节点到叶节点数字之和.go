/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
 *
 * https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/
 *
 * algorithms
 * Medium (69.39%)
 * Likes:    537
 * Dislikes: 0
 * Total Accepted:    165.6K
 * Total Submissions: 238.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
 *
 *
 * 每条从根节点到叶节点的路径都代表一个数字：
 *
 *
 * 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
 *
 *
 * 计算从根节点到叶节点生成的 所有数字之和 。
 *
 * 叶节点 是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3]
 * 输出：25
 * 解释：
 * 从根到叶子节点路径 1->2 代表数字 12
 * 从根到叶子节点路径 1->3 代表数字 13
 * 因此，数字总和 = 12 + 13 = 25
 *
 * 示例 2：
 *
 *
 * 输入：root = [4,9,0,5,1]
 * 输出：1026
 * 解释：
 * 从根到叶子节点路径 4->9->5 代表数字 495
 * 从根到叶子节点路径 4->9->1 代表数字 491
 * 从根到叶子节点路径 4->0 代表数字 40
 * 因此，数字总和 = 495 + 491 + 40 = 1026
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 1000] 内
 * 0
 * 树的深度不超过 10
 *
 *
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
func sumNumbers(root *TreeNode) int {
	res := 0
	res = dfs(root, 0)
	return res
}

func dfs(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	}
	if root.Left != nil && root.Right != nil {
		return dfs(root.Left, num) + dfs(root.Right, num)
	}
	if root.Left != nil {
		return dfs(root.Left, num)
	}
	return dfs(root.Right, num)
}

func sumNumbers2(root *TreeNode) int {
	res := 0
	res = dfs2(root, 0)
	return res
}

func dfs2(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	}

	return dfs2(root.Left, num) + dfs2(root.Right, num)
}

// @lc code=end
