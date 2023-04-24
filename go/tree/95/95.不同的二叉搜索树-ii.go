/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (73.23%)
 * Likes:    1416
 * Dislikes: 0
 * Total Accepted:    165.1K
 * Total Submissions: 225.4K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 *
 *
 */
package jzoffer

import (
	"github.com/yiGmMk/leetcode/datastructure"
)

type TreeNode = datastructure.TreeNode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/* TODO:
 * 二叉搜索树的性质: 右>根>左
 * 通过回溯解决(根据点为i)则左为(1,....i-1),右为(i+1...n)
 */
func generateTrees(n int) (res []*TreeNode) {
	if n == 0 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) (res []*TreeNode) {
	if start > end {
		res = append(res, nil)
		return
	}
	for i := start; i <= end; i++ {
		// 可行的左/右子树
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, l := range left {
			for _, r := range right {
				cur := &TreeNode{Val: i}
				cur.Left = l
				cur.Right = r
				res = append(res, cur)
			}
		}
	}
	return
}

// @lc code=end
