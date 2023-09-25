/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 *
 * https://leetcode.cn/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (42.83%)
 * Likes:    764
 * Dislikes: 0
 * Total Accepted:    136.7K
 * Total Submissions: 319.3K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
 *
 * 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回
 * true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,4,5]
 * 输出：true
 * 解释：任何 i < j < k 的三元组都满足题意
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,4,3,2,1]
 * 输出：false
 * 解释：不存在满足题意的三元组
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,1,5,0,4,6]
 * 输出：true
 * 解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^5
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
 *
 */
package jzoffer

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	n := len(nums)
	leftMin, rightMax := make([]int, n), make([]int, n)
	leftMin[0], rightMax[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}
	for i := n - 2; i > 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}
	for i := 1; i <= n-2; i++ {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end
