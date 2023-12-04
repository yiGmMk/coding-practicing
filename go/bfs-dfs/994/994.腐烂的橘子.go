/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 *
 * https://leetcode.cn/problems/rotting-oranges/description/
 *
 * algorithms
 * Medium (51.17%)
 * Likes:    781
 * Dislikes: 0
 * Total Accepted:    135.2K
 * Total Submissions: 264.2K
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]'
 *
 * 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
 *
 *
 * 值 0 代表空单元格；
 * 值 1 代表新鲜橘子；
 * 值 2 代表腐烂的橘子。
 *
 *
 * 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
 *
 * 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
 * 输出：-1
 * 解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[0,2]]
 * 输出：0
 * 解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10
 * grid[i][j] 仅为 0、1 或 2
 *
 *
 */
package jzoffer

// @lc code=start
func orangesRotting(grid [][]int) int {
	neighbours := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var (
		queue = [][]int{}
		good  int
	)

	for i, row := range grid {
		for j, v := range row {
			if v == 2 {
				queue = append(queue, []int{i, j})
			}
			if v == 1 {
				good++
			}
		}
	}
	if len(queue) == 0 {
		if good == 0 {
			return 0
		}
		return -1
	}

	check := func() bool {
		for _, row := range grid {
			for _, v := range row {
				if v == 1 {
					return false
				}
			}
		}
		return true
	}

	if check() {
		return 0
	}

	for res := 1; len(queue) > 0; res++ {
		tmp := queue
		queue = nil
		for _, v := range tmp {
			i, j := v[0], v[1]
			for _, nei := range neighbours {
				ni, nj := i+nei[0], j+nei[1]
				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					queue = append(queue, []int{ni, nj})
				}
			}
		}
		if check() {
			return res
		}
	}

	return -1
}

// @lc code=end
