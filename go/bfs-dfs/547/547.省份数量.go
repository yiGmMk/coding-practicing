/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 *
 * https://leetcode.cn/problems/number-of-provinces/description/
 *
 * algorithms
 * Medium (62.15%)
 * Likes:    1072
 * Dislikes: 0
 * Total Accepted:    291.9K
 * Total Submissions: 469.8K
 * Testcase Example:  '[[1,1,0],[1,1,0],[0,0,1]]'
 *
 *
 *
 * 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c
 * 间接相连。
 *
 * 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
 *
 * 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
 * isConnected[i][j] = 0 表示二者不直接相连。
 *
 * 返回矩阵中 省份 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * n == isConnected.length
 * n == isConnected[i].length
 * isConnected[i][j] 为 1 或 0
 * isConnected[i][i] == 1
 * isConnected[i][j] == isConnected[j][i]
 *
 *
 *
 *
 */
package jzoffer

// @lc code=start
func findCircleNum(isConnected [][]int) (res int) {
	// 1 1 0
	// 1 1 0
	// 0 0 1
	n := len(isConnected)
	visited := map[int]bool{}
	bfs := func(idx int) (num int) {
		if visited[idx] {
			return
		}
		visited[idx] = true
		queue := []int{idx}
		num++
		for len(queue) > 0 {
			top := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			row := isConnected[top] // row[j] == 1 表示城市 j 与城市 top 直接相连
			for j, v := range row {
				if v == 1 && !visited[j] {
					visited[j] = true
					queue = append(queue, j)
					num++
				}
			}

		}
		return
	}

	for i := 0; i < n; i++ {
		if bfs(i) > 0 {
			res++
		}
	}
	return
}

// @lc code=end
