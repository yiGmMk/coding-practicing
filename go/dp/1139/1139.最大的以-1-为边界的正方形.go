/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 *
 * https://leetcode.cn/problems/largest-1-bordered-square/description/
 *
 * algorithms
 * Medium (49.51%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    15.7K
 * Total Submissions: 30.3K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回
 * 0。
 *
 *
 *
 * 示例 1：
 *
 * 输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：9
 *
 *
 * 示例 2：
 *
 * 输入：grid = [[1,1,0,0]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= grid.length <= 100
 * 1 <= grid[0].length <= 100
 * grid[i][j] 为 0 或 1
 *
 *
 */
package jzoffer

// @lc code=start

// TODO

// ** 请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量**
func largest1BorderedSquare(grid [][]int) int {
	// 计算每个点为起点的上下左右各有多少连续的1
	m, n := len(grid), len(grid[0])
	left, up := make([][]int, m+1), make([][]int, m+1)
	//求出以(x,y)为起点4个反向上连续1的个数,枚举边长l就可以求出以(x,y)为右下角顶点的最大正方形
	for i := range left {
		left[i] = make([]int, n+1)
		up[i] = make([]int, n+1)
	}

	maxBoard := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] != 1 {
				continue
			}
			left[i][j] = left[i][j-1] + 1
			up[i][j] = up[i-1][j] + 1

			board := min(left[i][j], up[i][j])
			for ; left[i-board+1][j] < board || up[i][j-board+1] < board; board-- {
			}
			maxBoard = max(maxBoard, board)
		}
	}

	return maxBoard * maxBoard
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
