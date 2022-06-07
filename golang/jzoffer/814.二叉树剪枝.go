/*
 * @Author: yiGmMk marvelousme@163.com
 * @Date: 2022-06-07 14:43:50
 * @LastEditors: yiGmMk marvelousme@163.com
 * @LastEditTime: 2022-06-07 14:43:51
 * @FilePath: /go-tool/home/admin/code/coding-practicing/golang/jzoffer/814.二叉树剪枝.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @lc app=leetcode.cn id=814 lang=golang
 *
 * [814] 二叉树剪枝
 *
 * https://leetcode.cn/problems/binary-tree-pruning/description/
 *
 * algorithms
 * Medium (70.25%)
 * Likes:    228
 * Dislikes: 0
 * Total Accepted:    28.9K
 * Total Submissions: 41.1K
 * Testcase Example:  '[1,null,0,0,1]'
 *
 * 给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
 *
 * 返回移除了所有不包含 1 的子树的原二叉树。
 *
 * 节点 node 的子树为 node 本身加上所有 node 的后代。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,0,0,1]
 * 输出：[1,null,0,null,1]
 * 解释：
 * 只有红色节点满足条件“所有不包含 1 的子树”。 右图为返回的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,0,1,0,0,0,1]
 * 输出：[1,null,1,null,1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,1,0,1,1,0,1,0]
 * 输出：[1,1,0,1,1,null,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 200] 内
 * Node.val 为 0 或 1
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

func pruneTree(root *TreeNode) *TreeNode {

	return root
}

// @lc code=end
