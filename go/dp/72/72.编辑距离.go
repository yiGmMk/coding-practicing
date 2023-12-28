/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Medium (62.80%)
 * Likes:    3243
 * Dislikes: 0
 * Total Accepted:    426K
 * Total Submissions: 678.4K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 和 word2 由小写英文字母组成
 *
 *
 */
package jzoffer

// @lc code=start
// https://leetcode.cn/problems/edit-distance/solutions/188223/bian-ji-ju-chi-by-leetcode-solution/?envType=study-plan-v2&envId=leetcode-75
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	if n*m == 0 {
		return n + m
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown += 1
			}
			dp[i][j] = min(left, down, leftDown)
		}
	}
	return dp[n][m]
}

// @lc code=end
