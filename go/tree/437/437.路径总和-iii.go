/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 *
 * https://leetcode.cn/problems/path-sum-iii/description/
 *
 * algorithms
 * Medium (49.24%)
 * Likes:    1747
 * Dislikes: 0
 * Total Accepted:    255.9K
 * Total Submissions: 521.2K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
 *
 * 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
 * 输出：3
 * 解释：和等于 8 的路径有 3 条，如图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：3
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [0,1000]
 * -10^9
 * -1000
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
// dfs 时间复杂度O(N^2)
func pathSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	var (
		rootSum func(r *TreeNode, left int) int
	)
	rootSum = func(r *TreeNode, left int) int {
		if r == nil {
			return 0
		}
		val := 0
		if left-r.Val == 0 {
			val++
		}
		return val + rootSum(r.Left, left-r.Val) + rootSum(r.Right, left-r.Val)
	}
	// 以root开始的路径
	res += rootSum(root, targetSum)
	// 以root.Left为起点的路径
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return
}

// 前缀和
func pathSumPrefixSum(root *TreeNode, targetSum int) (res int) {
	var (
		dfs    func(r *TreeNode, sum int)
		preSum = make(map[int]int)
	)

	// 只有一个节点的路径也支持,所以需要设置前缀和为0的数量=1
	preSum[0] = 1
	dfs = func(r *TreeNode, sum int) {
		if r == nil {
			return
		}
		// 根节点到当前节点的前缀和 - targetSum
		// 如果存在于map中说明存在有路径可以删除使部分路径的和为tagetSum
		res += preSum[(sum+r.Val)-targetSum]
		preSum[sum+r.Val]++
		dfs(r.Left, sum+r.Val)
		dfs(r.Right, sum+r.Val)
		preSum[sum+r.Val]--
		return
	}
	dfs(root, 0)
	return
}

// @lc code=end
