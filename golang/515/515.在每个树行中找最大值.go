/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (65.62%)
 * Likes:    194
 * Dislikes: 0
 * Total Accepted:    63.3K
 * Total Submissions: 96.4K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入: root = [1,3,2,5,3,null,9]
 * 输出: [1,3,9]
 *
 *
 * 示例2：
 *
 *
 * 输入: root = [1,2,3]
 * 输出: [1,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 二叉树的节点个数的范围是 [0,10^4]
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
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

import (
	"container/list"

	"github.com/yiGmMk/leetcode/golang/util"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode util.TreeNode

func largestValues(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	level := list.New()
	for n := level.Len(); n > 0; {
		max := level.Front().Value.(*TreeNode).Val
		for i := 0; i < n; i++ {
			val := level.Front()
			if val == nil {
				continue
			}
			max = Max(max, val.Value.(*TreeNode).Val)
			left := val.Value.(*TreeNode).Left
			if left != nil {
				level.PushBack(left)
			}
			right := val.Value.(*TreeNode).Right
			if right != nil {
				level.PushBack(right)
			}
			level.Remove(val)
		}
		out = append(out, max)
	}
	return out
}

// @lc code=end
