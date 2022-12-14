/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.34%)
 * Likes:    2025
 * Dislikes: 0
 * Total Accepted:    688.7K
 * Total Submissions: 1.6M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
 *
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 *
 * 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 *
 * 示例 3：
 *
 *
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums 是一个非递减数组
 * -10^9 <= target <= 10^9
 *
 *
 */
package jzoffer

import (
	"sort"

	binarysearch "github.com/yiGmMk/leetcode/go/binary-search"
)

// @lc code=start
func searchRange0(nums []int, target int) (res []int) {
	i := sort.SearchInts(nums, target)
	if i >= len(nums) || nums[i] != target {
		res = append(res, -1)
	} else {
		res = append(res, i)
	}
	i = sort.SearchInts(nums, target+1)
	if i > len(nums) || i < 1 || nums[i-1] != target {
		res = append(res, -1)
	} else {
		res = append(res, i-1)
	}
	return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/zai-pai-xu-shu-zu-zhong-cha-zhao-yuan-su-de-di-3-4/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func searchRange(nums []int, target int) []int {
	left := binarysearch.LowerBound(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := binarysearch.UpperBound(nums, target) - 1
	return []int{left, right}
}

// @lc code=end
