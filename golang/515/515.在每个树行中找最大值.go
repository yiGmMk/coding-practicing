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
//TODO  提交时还不能带package包
//package jzoffer

import (
	"container/list"
	"math"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 使用go的list实现
func largestValuesUsingList(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	level := list.New()
	level.PushBack(root)
	for level.Len() > 0 {
		nl := level.Len()
		max := math.MinInt32
		for i := 0; i < nl; i++ {
			node := level.Front()
			level.Remove(node)
			val := node.Value.(*TreeNode)
			if val == nil {
				continue
			}
			max = Max(max, val.Val)
			left := val.Left
			if left != nil {
				level.PushBack(left)
			}
			right := val.Right
			if right != nil {
				level.PushBack(right)
			}
		}
		out = append(out, max)
	}
	return out
}

// 使用切片实现
func largestValues(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	level := []*TreeNode{root}
	levelNum := len(level)
	for levelNum > 0 {
		max := math.MinInt32
		for i := 0; i < levelNum; i++ {
			node := level[0]
			level = level[1:]
			if node == nil {
				continue
			}
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		out = append(out, max)
		levelNum = len(level)
	}
	return out
}

// @lc code=end
