package leetcode

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 *
 * https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (51.08%)
 * Likes:    1033
 * Dislikes: 0
 * Total Accepted:    274.8K
 * Total Submissions: 536.4K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
 *
 *
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 5
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 20
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= n, m <= 300
 * -10^9 <= matrix[i][j] <= 10^9
 * 每行的所有元素从左到右升序排列
 * 每列的所有元素从上到下升序排列
 * -10^9 <= target <= 10^9
 *
 *
*/

// @lc code=start
// func searchMatrix(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return false
// 	}
// 	// v=matrix[y][x]
// 	for y := len(matrix) - 1; y >= 0; y-- {
// 		for x := 0; x < len(matrix[y]); x++ {
// 			if y > (len(matrix)-1) || x > (len(matrix[y])-1) {
// 				return false
// 			}
// 			v := matrix[y][x]
// 			if v == target {
// 				return true
// 			} else if v > target && y > 0 {
// 				y--
// 				x = -1 // 回到第一列,每次循环后x++会回到0
// 			}
// 		}
// 	}
// 	return false
// }

// func searchMatrix(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return false
// 	}

// 	// v=matrix[y][x]
// 	for y, row := range matrix {
// 		for x := len(row) - 1; x >= 0; x-- {
// 			if y > (len(matrix)-1) || x > (len(matrix[y])-1) {
// 				return false
// 			}
// 			val := matrix[y][x]
// 			if val == target {
// 				return true
// 			} else if val < target {
// 				y++
// 				x++
// 			}
// 		}
// 	}
// 	return false
// }

/*
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/search-a-2d-matrix-ii/solution/sou-suo-er-wei-ju-zhen-ii-by-leetcode-so-9hcx/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

// @lc code=end
