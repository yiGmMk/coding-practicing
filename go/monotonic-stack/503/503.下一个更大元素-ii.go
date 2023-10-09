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
func nextGreaterElementsMy(nums []int) []int {
	var (
		idx = -1
		n   = len(nums)
		// 环形数组,多一份数据用于模拟
		num   = append(nums, nums...)
		stack = []struct{ i, v int }{}
		res   = make([]int, len(nums))
	)
	// 默认没有下一个更大的值,有则在单调栈中更新
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	for _, v := range num {
		idx++
		if idx == n { //到数组尾部了,开始找前头有没有更大的值(环形)
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

// 实际不需要拷贝一份数据,多遍历一次就好了
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/next-greater-element-ii/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
