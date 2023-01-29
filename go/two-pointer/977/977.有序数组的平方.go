/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 *
 * https://leetcode.cn/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (68.66%)
 * Likes:    705
 * Dislikes: 0
 * Total Accepted:    461.8K
 * Total Submissions: 672.6K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 * 解释：平方后，数组变为 [16,1,0,9,100]
 * 排序后，数组变为 [0,1,9,16,100]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 10^4
 * -10^4
 * nums 已按 非递减顺序 排序
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 请你设计时间复杂度为 O(n) 的算法解决本问题
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
// 先按绝对值排序再计算
func sortedSquares1(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return abs(nums[i]) < abs(nums[j]) })

	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v * v
	}
	return res
}

// 直接修改原数组
func sortedSquares2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return abs(nums[i]) < abs(nums[j]) })

	for i, v := range nums {
		nums[i] = v * v
	}
	return nums
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// 先找最接近0值的数据,完两边遍历数组
func sortedSquares3(nums []int) []int {
	i := sort.SearchInts(nums, 0)
	l, r := i-1, i
	var res []int
	for l >= 0 || r < len(nums) {
		vl, vr := -1, -1
		if l >= 0 {
			vl = nums[l] * nums[l]
		}
		if r < len(nums) {
			vr = nums[r] * nums[r]
		}
		if vl >= 0 && vr >= 0 {
			if vl > vr {
				res = append(res, vr)
				r++
			} else {
				res = append(res, vl)
				l--
			}
			continue
		}
		if vl >= 0 {
			res = append(res, vl)
			l--
		}
		if vr >= 0 {
			res = append(res, vr)
			r++
		}
	}
	return res
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/squares-of-a-sorted-array/solution/you-xu-shu-zu-de-ping-fang-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}

// @lc code=end
