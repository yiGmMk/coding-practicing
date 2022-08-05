/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心下标
 *
 * https://leetcode.cn/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (49.73%)
 * Likes:    445
 * Dislikes: 0
 * Total Accepted:    209.1K
 * Total Submissions: 420K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * 给你一个整数数组 nums ，请计算数组的 中心下标 。
 *
 * 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
 *
 * 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
 *
 * 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1, 7, 3, 6, 5, 6]
 * 输出：3
 * 解释：
 * 中心下标是 3 。
 * 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
 * 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1, 2, 3]
 * 输出：-1
 * 解释：
 * 数组中不存在满足此条件的中心下标。
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2, 1, -1]
 * 输出：0
 * 解释：
 * 中心下标是 0 。
 * 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
 * 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * -1000 <= nums[i] <= 1000
 *
 *
 *
 *
 * 注意：本题与主站 1991
 * 题相同：https://leetcode-cn.com/problems/find-the-middle-index-in-array/
 *
 */
package jzoffer

// @lc code=start
func pivotIndex(nums []int) int {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}

	n := len(nums)
	// 1, 7, 3, 6, 5, 6
	// 0  1  8  11 17 22 28
	for i := 1; i <= n; i++ {
		// left == right
		if sums[i-1] == sums[n]-sums[i] {
			return i - 1
		}
	}
	return -1
}

// @lc code=end
