/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (48.58%)
 * Likes:    1270
 * Dislikes: 0
 * Total Accepted:    380.1K
 * Total Submissions: 782.9K
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
)

// @lc code=start
func minSubArrayLenTimeout(target int, nums []int) int {
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum >= target {
			return 1
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				res = min(res, j-i+1)
			}
		}
	}
	if res != math.MaxInt {
		return res
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func minSubArrayLen(target int, nums []int) int {
// 	res := math.MaxInt
// 	sum, sums := 0, make(map[int]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		sum += nums[i]
// 		sums[i] = sum
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		num := target + sums[i]
// 		bound := sort.SearchInts(nums, num)
// 		if bound < 0 {
// 			bound = -bound
// 		}
// 		if bound < len(nums) {
// 			res = min(res, bound-i)
// 		}
// 	}
// 	if res != math.MaxInt {
// 		return res
// 	}
// 	return 0
// }

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/2VG8Kg/solution/he-da-yu-deng-yu-target-de-zui-duan-zi-s-ixef/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
