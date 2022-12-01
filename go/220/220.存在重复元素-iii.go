/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode.cn/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (29.00%)
 * Likes:    625
 * Dislikes: 0
 * Total Accepted:    81.9K
 * Total Submissions: 281.9K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j])
 * ，同时又满足 abs(i - j)  。
 *
 * 如果存在则返回 true，不存在返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,1], k = 3, t = 0
 * 输出：true
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,0,1,1], k = 1, t = 2
 * 输出：true
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,5,9,1,5,9], k = 2, t = 3
 * 输出：false
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -2^31
 * 0
 * 0
 *
 *
 */

package jzoffer

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	out := false
	// abs(nums[i] - nums[j])<=t,同时又满足 abs(i - j)<=k  。
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && abs(i-j) <= k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
