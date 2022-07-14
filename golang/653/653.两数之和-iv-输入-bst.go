/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
 *
 * https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (63.17%)
 * Likes:    410
 * Dislikes: 0
 * Total Accepted:    94.5K
 * Total Submissions: 149.5K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * 给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], k = 9
 * 输出: true
 *
 *
 * 示例 2：
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], k = 28
 * 输出: false
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是  [1, 10^4].
 * -10^4 <= Node.val <= 10^4
 * root 为二叉搜索树
 * -10^5 <= k <= 10^5
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package jzoffer

import "github.com/yiGmMk/leetcode/golang/util"

type TreeNode = util.TreeNode

func findTarget(root *TreeNode, k int) bool {
	nodeMap := make(map[int]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		nodeMap[root.Val]++
	}
	dfs(root)
	for v1, v1num := range nodeMap {
		if _, ok := nodeMap[k-v1]; ok {
			if (v1 == k-v1 && v1num > 1) || v1 != k-v1 {
				return true
			}
		}
	}
	return false
}

// @lc code=end
