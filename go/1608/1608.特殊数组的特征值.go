/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 *
 * https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/description/
 *
 * algorithms
 * Easy (61.65%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    51.5K
 * Total Submissions: 83.5K
 * Testcase Example:  '[3,5]'
 *
 * 给你一个非负整数数组 nums 。如果存在一个数 x ，使得 nums 中恰好有 x 个元素 大于或者等于 x ，那么就称 nums 是一个 特殊数组
 * ，而 x 是该数组的 特征值 。
 *
 * 注意： x 不必 是 nums 的中的元素。
 *
 * 如果数组 nums 是一个 特殊数组 ，请返回它的特征值 x 。否则，返回 -1 。可以证明的是，如果 nums 是特殊数组，那么其特征值 x 是
 * 唯一的 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,5]
 * 输出：2
 * 解释：有 2 个元素（3 和 5）大于或等于 2 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [0,0]
 * 输出：-1
 * 解释：没有满足题目要求的特殊数组，故而也不存在特征值 x 。
 * 如果 x = 0，应该有 0 个元素 >= x，但实际有 2 个。
 * 如果 x = 1，应该有 1 个元素 >= x，但实际有 0 个。
 * 如果 x = 2，应该有 2 个元素 >= x，但实际有 0 个。
 * x 不能取更大的值，因为 nums 中只有两个元素。
 *
 * 示例 3：
 *
 * 输入：nums = [0,4,3,0,4]
 * 输出：3
 * 解释：有 3 个元素大于或等于 3 。
 *
 *
 * 示例 4：
 *
 * 输入：nums = [3,6,7,7,0]
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 1000
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
func specialArray1(nums []int) int {
	check := func(x int) int {
		v := 0
		for _, val := range nums {
			if val >= x {
				v++
			}
		}
		return v
	}
	n := sort.Search(len(nums), func(i int) bool {
		return check(i) < i
	})
	if check(n) == n {
		return n
	}
	if check(n-1) == (n - 1) {
		return n - 1
	}
	return -1
}

// TODO
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/solution/te-shu-shu-zu-de-te-zheng-zhi-by-leetcod-9wfo/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func specialArray(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) //降序排序
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

// @lc code=end
