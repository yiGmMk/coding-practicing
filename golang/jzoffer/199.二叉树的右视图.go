/*
 * @Author: yiGmMk marvelousme@163.com
 * @Date: 2022-06-07 14:43:55
 * @LastEditors: yiGmMk marvelousme@163.com
 * @LastEditTime: 2022-06-07 14:43:56
 * @FilePath: /go-tool/home/admin/code/coding-practicing/golang/jzoffer/199.二叉树的右视图.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode.cn/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (65.54%)
 * Likes:    699
 * Dislikes: 0
 * Total Accepted:    211.4K
 * Total Submissions: 322.6K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1,3,4]
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,null,3]
 * 输出: [1,3]
 *
 *
 * 示例 3:
 *
 *
 * 输入: []
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [0,100]
 * -100
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

// import "github.com/yiGmMk/leetcode/golang/util"

// type TreeNode = util.TreeNode

// func rightSideView(root *TreeNode) []int {
// 	out := []int{}
// 	if root == nil {
// 		return out
// 	}
// 	queue := []*TreeNode{root}
// 	qL := len(queue)
// 	for qL > 0 {
// 		n := len(queue)
// 		for i := 0; i < n; i++ {
// 			node := queue[0]
// 			queue = queue[1:]
// 			if node == nil {
// 				continue
// 			}
// 			if i == n-1 {
// 				out = append(out, node.Val)
// 			}
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		qL = len(queue)
// 	}
// 	return out
// }

// @lc code=end
