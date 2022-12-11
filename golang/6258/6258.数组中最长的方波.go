/*
 * @lc app=leetcode.cn id=6258 lang=golang
 *
 * [6258] 数组中最长的方波
 *
 * https://leetcode.cn/problems/longest-square-streak-in-an-array/description/
 *
 * algorithms
 * Medium (39.68%)
 * Likes:    3
 * Dislikes: 0
 * Total Accepted:    4.1K
 * Total Submissions: 10.2K
 * Testcase Example:  '[4,3,6,16,8,2]'
 *
 * 给你一个整数数组 nums 。如果 nums 的子序列满足下述条件，则认为该子序列是一个 方波 ：
 *
 *
 * 子序列的长度至少为 2 ，并且
 * 将子序列从小到大排序 之后 ，除第一个元素外，每个元素都是前一个元素的 平方 。
 *
 *
 * 返回 nums 中 最长方波 的长度，如果不存在 方波 则返回 -1 。
 *
 * 子序列 也是一个数组，可以由另一个数组删除一些或不删除元素且不改变剩余元素的顺序得到。
 *
 *
 *
 * 示例 1 ：
 *
 * 输入：nums = [4,3,6,16,8,2]
 * 输出：3
 * 解释：选出子序列 [4,16,2] 。排序后，得到 [2,4,16] 。
 * - 4 = 2 * 2.
 * - 16 = 4 * 4.
 * 因此，[4,16,2] 是一个方波.
 * 可以证明长度为 4 的子序列都不是方波。
 *
 *
 * 示例 2 ：
 *
 * 输入：nums = [2,3,5,6,7]
 * 输出：-1
 * 解释：nums 不存在方波，所以返回 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * 2 <= nums[i] <= 10^5
 *
 *
 */
package jzoffer

// @lc code=start
// 暴力搜索
func longestSquareStreakBruceForce(nums []int) (res int) {
	vs := make(map[int]bool, len(nums))
	for _, v := range nums {
		vs[v] = true
	}
	for _, v := range nums {
		cnt := 1 //v 肯定在vs里,跳过,直接查v*v
		for x := v * v; vs[x]; x *= x {
			cnt++
		}
		if cnt > res {
			res = cnt
		}
	}
	if res > 1 {
		return
	}
	return -1
}

// 使用map缓存中间结果
func longestSquareStreak(nums []int) (res int) {
	vs := make(map[int]bool, len(nums))
	for _, v := range nums {
		vs[v] = true
	}
	dp := make(map[int]int, len(nums))
	var dfs func(v int) int
	dfs = func(v int) int {
		if !vs[v] {
			return 0
		}
		if cnt, ok := dp[v]; ok {
			return cnt
		}
		dp[v] = 1 + dfs(v*v)
		return dp[v]
	}
	for _, v := range nums {
		cnt := dfs(v)
		if cnt > res {
			res = cnt
		}
	}
	if res > 1 {
		return
	}
	return -1
}

// @lc code=end
