/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 *
 * https://leetcode.cn/problems/binary-search/description/
 *
 * algorithms
 * Easy (54.55%)
 * Likes:    1105
 * Dislikes: 0
 * Total Accepted:    843.1K
 * Total Submissions: 1.5M
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的
 * target，如果目标值存在返回下标，否则返回 -1。
 *
 *
 * 示例 1:
 *
 * 输入: nums = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 出现在 nums 中并且下标为 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不存在 nums 中因此返回 -1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 你可以假设 nums 中的所有元素是不重复的。
 * n 将在 [1, 10000]之间。
 * nums 的每个元素都将在 [-9999, 9999]之间。
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
func search1(nums []int, target int) int {
	i := sort.SearchInts(nums, target)
	if i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}

func search(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	if l >= 0 && l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

// @lc code=end
