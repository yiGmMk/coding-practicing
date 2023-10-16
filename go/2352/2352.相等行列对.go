/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] 相等行列对
 *
 * https://leetcode.cn/problems/equal-row-and-column-pairs/description/
 *
 * algorithms
 * Medium (73.35%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    39.7K
 * Total Submissions: 54.2K
 * Testcase Example:  '[[3,2,1],[1,7,6],[2,7,7]]'
 *
 * 给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。
 *
 * 如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid = [[3,2,1],[1,7,6],[2,7,7]]
 * 输出：1
 * 解释：存在一对相等行列对：
 * - (第 2 行，第 1 列)：[2,7,7]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
 * 输出：3
 * 解释：存在三对相等行列对：
 * - (第 0 行，第 0 列)：[3,1,2,2]
 * - (第 2 行, 第 2 列)：[2,4,2,2]
 * - (第 3 行, 第 2 列)：[2,4,2,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length == grid[i].length
 * 1 <= n <= 200
 * 1 <= grid[i][j] <= 10^5
 *
 *
 */
package jzoffer

import (
	"strconv"
	"strings"
)

// @lc code=start
func trans(v ...int) string {
	b := strings.Builder{}
	n := len(v)
	for i, val := range v {
		b.WriteString(strconv.Itoa(val))
		if i != n-1 {
			b.WriteString(",")
		}
	}
	return b.String()
}

// 11 1
// 1  11
func equalPairs(grid [][]int) (res int) {
	row := make(map[string]int, len(grid))
	for _, v := range grid {
		val := trans(v...)
		row[val] = row[val] + 1
	}
	n, m := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		col := []int{}
		for j := 0; j < n; j++ {
			col = append(col, grid[j][i])
		}
		val := trans(col...)
		if _, ok := row[val]; ok {
			res += row[val]
		}
	}
	return
}

// @lc code=end
