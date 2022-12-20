/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (47.56%)
 * Likes:    1489
 * Dislikes: 0
 * Total Accepted:    460.3K
 * Total Submissions: 968.2K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。
 *
 * 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr]
 * ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：target = 4, nums = [1,4,4]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
 *
 *
 */
package jzoffer

import (
	"math"
	"sort"
)

// @lc code=start
// 前缀和+二分查找,时间复杂度 O(n log(n))
func minSubArrayLen1(target int, nums []int) int {
	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] = nums[i-1] + sum[i-1]
	}

	min := math.MaxInt
	for i := 0; i < len(sum); i++ {
		j := sort.SearchInts(sum, sum[i]+target)
		if j < len(sum) {
			if j-i < min {
				min = j - i
			}
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}

// TODO
// 每次确定子数组的开始下标，然后得到长度最小的子数组，因此时间复杂度较高。
// 为了降低时间复杂度，可以使用滑动窗口的方法
func minSubArrayLen(target int, nums []int) int {
	n, sum := len(nums), 0
	l, r, min := 0, 0, math.MaxInt
	for r < n {
		sum += nums[r]
		for sum >= target {
			sum -= nums[l]
			if r-l+1 < min {
				min = r - l + 1
			}
			l++
		}
		r++
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}

// @lc code=end
