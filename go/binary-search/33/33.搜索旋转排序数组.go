/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.82%)
 * Likes:    2415
 * Dislikes: 0
 * Total Accepted:    658.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 *
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始
 * 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
 *
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 *
 * 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * nums 中的每个值都 独一无二
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 <= target <= 10^4
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
func search(nums []int, target int) (res int) {
	// 找到最小值,跳出循环后l,r不相等,l+1=r
	// findMin := func(nums []int) int {
	// 	l, r := -1, len(nums)-1 // 开区间 (-1, n-1)
	// 	for l+1 < r { // 开区间不为空
	// 		mid := int(uint(l+r)) >> 1
	// 		if nums[mid] < nums[len(nums)-1] { // mid在后半部
	// 			r = mid
	// 		} else {
	// 			l = mid
	// 		}
	// 	}
	// 	return r
	// }
	// 跳出循环后l==r
	findMin := func(nums []int) int {
		l, r := 0, len(nums)-1
		for l < r {
			mid := int(uint(l+r) >> 1)
			if nums[mid] >= nums[r] { // mid在前半部
				l = mid + 1
			} else {
				r = mid
			}
		}
		return r
	}
	min := findMin(nums)
	i := -1
	// 分段搜索
	if target > nums[len(nums)-1] {
		i = sort.SearchInts(nums[:min], target)
		if i < len(nums[:min]) && nums[i] == target {
			return i
		}
	} else {
		i = sort.SearchInts(nums[min:], target)
		if i < len(nums[min:]) && nums[i+min] == target {
			return i + min
		}
	}
	return -1
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 退出循环时一定有 low==high
	return nums[l] // nums[high] 也行
}

// @lc code=end
