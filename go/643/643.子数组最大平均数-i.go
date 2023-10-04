/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 *
 * https://leetcode.cn/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (43.14%)
 * Likes:    324
 * Dislikes: 0
 * Total Accepted:    119.9K
 * Total Submissions: 277.9K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 * 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
 *
 * 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
 *
 * 任何误差小于 10^-5 的答案都将被视为正确答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,12,-5,-6,50,3], k = 4
 * 输出：12.75
 * 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5], k = 1
 * 输出：5.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= k <= n <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
package jzoffer

import "math"

// @lc code=start
func findMaxAverage1(nums []int, k int) (res float64) {
	res = float64(math.MinInt)
	sum, i := float64(0), 0
	for j, v := range nums {
		sum += float64(v)
		i++
		if i >= k {
			if i > k {
				sum -= float64(nums[j-k])
			}
			res = max(sum/float64(k), res)
		}
	}

	return
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func findMaxAverage(nums []int, k int) (res float64) {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	res = float64(math.MinInt)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		if i >= k-1 {
			sum := nums[i]
			if i > k-1 {
				sum -= nums[i-k]
			}
			res = max(float64(sum)/float64(k), res)
		}
	}
	return
}

// @lc code=end
