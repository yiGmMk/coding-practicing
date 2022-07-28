/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode.cn/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (72.11%)
 * Likes:    893
 * Dislikes: 0
 * Total Accepted:    129.8K
 * Total Submissions: 180K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,2,3,2]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,1,0,1,99]
 * 输出：99
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 * nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 *
 *
 *
 *
 * 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 */
package jzoffer

import "sort"

// 1 1 1 10 10 10 11
// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	var res = nums[0]
	size := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			size++
			continue
		}
		// not equal
		if size == 1 {
			return res
		}
		size = 1
		res = nums[i]
	}
	return res
}

// @lc code=end
