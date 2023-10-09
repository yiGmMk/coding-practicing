/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode.cn/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (67.02%)
 * Likes:    872
 * Dislikes: 0
 * Total Accepted:    215.7K
 * Total Submissions: 321.8K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的
 * 下一个更大元素 。
 *
 * 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
 * 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数；
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4,3]
 * 输出: [2,3,4,-1,4]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */
package jzoffer

// @lc code=start
func nextGreaterElements(nums []int) []int {
	var (
		idx   = -1
		n     = len(nums)
		num   = append(nums, nums...)
		stack = []struct{ i, v int }{}
		res   = make([]int, len(nums))
	)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for _, v := range num {
		idx++
		if idx == n {
			idx = 0
		}
		for len(stack) > 0 && v > stack[len(stack)-1].v {
			top := stack[len(stack)-1]
			res[top.i] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, struct {
			i int
			v int
		}{i: idx, v: v})
	}
	return res
}

// @lc code=end
