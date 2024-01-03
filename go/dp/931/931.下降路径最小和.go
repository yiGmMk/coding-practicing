/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 *
 * https://leetcode.cn/problems/minimum-falling-path-sum/description/
 *
 * algorithms
 * Medium (67.54%)
 * Likes:    327
 * Dislikes: 0
 * Total Accepted:    94.3K
 * Total Submissions: 139.6K
 * Testcase Example:  '[[2,1,3],[6,5,4],[7,8,9]]'
 *
 * 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
 *
 * 下降路径
 * 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置
 * (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1)
 * 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
 * 输出：13
 * 解释：如图所示，为和最小的两条下降路径
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：matrix = [[-19,57],[-40,-5]]
 * 输出：-59
 * 解释：如图所示，为和最小的下降路径
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 100
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
package jzoffer

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	dp[0][0] = matrix[0][0]
	m, n := len(matrix), len(matrix[0])

	// dp[i][j]=min(dp[i-1][j-1],dp[i-1][j],dp[i-1][j+1])+matrix[i][j]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0:
				dp[i][j] = matrix[i][j]
			case j == 0:
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			case j == n-1:
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + matrix[i][j]
			default:
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			}
		}
	}
	res := dp[m-1][0]
	for j := 1; j < n; j++ {
		res = min(res, dp[m-1][j])
	}
	return res
}

// @lc code=end
