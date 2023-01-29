/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 *
 * https://leetcode.cn/problems/sum-of-square-numbers/description/
 *
 * algorithms
 * Medium (38.52%)
 * Likes:    409
 * Dislikes: 0
 * Total Accepted:    125.5K
 * Total Submissions: 325.9K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：c = 5
 * 输出：true
 * 解释：1 * 1 + 2 * 2 = 5
 *
 *
 * 示例 2：
 *
 *
 * 输入：c = 3
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= c <= 2^31 - 1
 *
 *
 */
package jzoffer

import "math"

// @lc code=start
func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		v := l*l + r*r
		if v > c {
			r--
		} else if v == c {
			return true
		} else {
			l++
		}
	}
	return false
}

// @lc code=end
