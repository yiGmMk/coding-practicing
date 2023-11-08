/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
 *
 * https://leetcode.cn/problems/longest-univalue-path/description/
 *
 * algorithms
 * Medium (47.77%)
 * Likes:    782
 * Dislikes: 0
 * Total Accepted:    84.8K
 * Total Submissions: 177.3K
 * Testcase Example:  '[5,4,5,1,1,null,5]'
 *
 * 给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
 *
 * 两个节点之间的路径长度 由它们之间的边数表示。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：root = [5,4,5,1,1,5]
 * 输出：2
 *
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入：root = [1,4,5,4,4,5]
 * 输出：2
 *
 *
 *
 *
 * 提示:
 *
 *
 * 树的节点数的范围是 [0, 10^4]
 * -1000 <= Node.val <= 1000
 * 树的深度将不超过 1000
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
func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		if !(node.Left != nil && node.Left.Val == node.Val) {
			left = 0
		}
		if !(node.Right != nil && node.Right.Val == node.Val) {
			right = 0
		}
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/longest-univalue-path/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
