/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 *
 * https://leetcode.cn/problems/n-th-tribonacci-number/description/
 *
 * algorithms
 * Easy (60.94%)
 * Likes:    291
 * Dislikes: 0
 * Total Accepted:    181.2K
 * Total Submissions: 297.4K
 * Testcase Example:  '4'
 *
 * 泰波那契序列 Tn 定义如下：
 *
 * T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
 *
 * 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 4
 * 输出：4
 * 解释：
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 *
 *
 * 示例 2：
 *
 * 输入：n = 25
 * 输出：1389537
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 37
 * 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
 *
 *
 */
package jzoffer

// @lc code=start
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	n0, n1, n2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		n3 := n0 + n1 + n2
		n0, n1, n2 = n1, n2, n3
	}
	return n2
}

// @lc code=end
