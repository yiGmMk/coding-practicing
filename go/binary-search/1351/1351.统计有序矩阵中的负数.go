/*
 * @lc app=leetcode.cn id=1351 lang=golang
 *
 * [1351] 统计有序矩阵中的负数
 *
 * https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/description/
 *
 * algorithms
 * Easy (74.61%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    46.8K
 * Total Submissions: 62.7K
 * Testcase Example:  '[[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]'
 *
 * 给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列。 请你统计并返回 grid 中 负数 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
 * 输出：8
 * 解释：矩阵中共有 8 个负数。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[3,2],[1,0]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 100
 * -100 <= grid[i][j] <= 100
 *
 *
 *
 *
 * 进阶：你可以设计一个时间复杂度为 O(n + m) 的解决方案吗？
 *
 */
package jzoffer

import "sort"

// @lc code=start

func countNegatives1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][0] < 0 {
		return m * n
	}
	if grid[m-1][n-1] >= 0 {
		return 0
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			// fmt.Println(grid[i][j])
			if grid[i][j] >= 0 {
				break
			} else {
				res++
			}
		}
	}
	return res
}

func countNegatives(grid [][]int) int {
	res := 0
	for _, row := range grid {
		n := sort.Search(len(row), func(i int) bool { return row[i] < 0 })
		if n < len(row) {
			res += len(row) - n
		}
	}
	return res
}

// @lc code=end
