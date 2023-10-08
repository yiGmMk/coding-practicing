/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode.cn/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (68.73%)
 * Likes:    1626
 * Dislikes: 0
 * Total Accepted:    471.3K
 * Total Submissions: 685.5K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i
 * 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: temperatures = [73,74,75,71,69,72,76,73]
 * 输出: [1,1,4,2,1,1,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: temperatures = [30,40,50,60]
 * 输出: [1,1,1,0]
 *
 *
 * 示例 3:
 *
 *
 * 输入: temperatures = [30,60,90]
 * 输出: [1,1,0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= temperatures.length <= 10^5
 * 30 <= temperatures[i] <= 100
 *
 *
 */
package jzoffer

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	var (
		idx   = -1
		stack = [][2]int{}
		res   = make([]int, len(temperatures))
	)
	for _, v := range temperatures {
		idx++
		for len(stack) > 0 && v > stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			res[top[0]] = idx - top[0]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{idx, v})
	}
	return res
}

// @lc code=end
