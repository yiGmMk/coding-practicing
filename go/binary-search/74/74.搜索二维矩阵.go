/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode.cn/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (48.36%)
 * Likes:    745
 * Dislikes: 0
 * Total Accepted:    284K
 * Total Submissions: 587.2K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
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
 * 1
 * -10^4
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
// 和2的思路是一致的,写的太繁琐了,好几个分支可以合并
func searchMatrix(matrix [][]int, target int) bool {
	col := []int{}
	for i := range matrix {
		col = append(col, matrix[i][0])
	}
	row := sort.SearchInts(col, target)
	if row >= len(matrix) { //最后一行
		v := sort.SearchInts(matrix[len(matrix)-1], target)
		if v >= len(matrix[0]) {
			return false
		}
		return matrix[len(matrix)-1][v] == target
	}
	if matrix[row][0] == target { // 第一列
		return true
	}
	if row <= 0 { // 不存在
		return false
	}
	// 比第一列的数大,在该行找
	v := sort.SearchInts(matrix[row-1], target)
	if v >= len(matrix[0]) {
		return false
	}
	return matrix[row-1][v] == target
}

func searchMatrix1(matrix [][]int, target int) bool {
	col := []int{}
	for i := range matrix {
		col = append(col, matrix[i][0])
	}
	row := sort.SearchInts(col, target) // row >= target
	if row < len(matrix) && matrix[row][0] == target {
		return true
	}
	if row <= 0 { // 不存在
		return false
	}
	// 比第一列的数大,在该行找
	v := sort.SearchInts(matrix[row-1], target)
	if v >= len(matrix[0]) {
		return false
	}
	return matrix[row-1][v] == target
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/search-a-2d-matrix/solution/sou-suo-er-wei-ju-zhen-by-leetcode-solut-vxui/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// 两次搜索,先搜所在行,再找列,找<=target的那一列方便很多
func searchMatrix2(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix),
		func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

// 一次搜索
func searchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
	return i < m*n && matrix[i/n][i%n] == target
}

// @lc code=end
