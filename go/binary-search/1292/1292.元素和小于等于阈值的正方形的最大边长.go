/*
 * @lc app=leetcode.cn id=1292 lang=golang
 *
 * [1292] 元素和小于等于阈值的正方形的最大边长
 *
 * https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/description/
 *
 * algorithms
 * Medium (49.77%)
 * Likes:    105
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 20.5K
 * Testcase Example:  '[[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]]\n4'
 *
 * 给你一个大小为 m x n 的矩阵 mat 和一个整数阈值 threshold。
 *
 * 请你返回元素总和小于或等于阈值的正方形区域的最大边长；如果没有这样的正方形区域，则返回 0 。
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
 * 输出：2
 * 解释：总和小于或等于 4 的正方形的最大边长为 2，如图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]],
 * threshold = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1 <= m, n <= 300
 * 0 <= mat[i][j] <= 10^4
 * 0 <= threshold <= 10^5^
 *
 *
 */

package jzoffer

// TODO
// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 时间复杂度,O(MN*log min(M,N))
func maxSideLength1(mat [][]int, threshold int) (ans int) {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	check := func(x1, y1, x2, y2 int) bool {
		return sum[x2][y2]-sum[x1-1][y2]-sum[x2][y1-1]+sum[x1-1][y1-1] <= threshold
	}
	l, r, ans := 1, min(m, n), 0
	for l <= r {
		mid := int(uint(l+r) >> 1)
		find := false
		for i := 1; i <= m-mid+1; i++ {
			for j := 1; j <= n-mid+1; j++ {
				if check(i, j, i+mid-1, j+mid-1) {
					find = true
				}
			}
		}
		if find {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 舍弃二分,直接枚举+优化
func maxSideLength(mat [][]int, threshold int) (ans int) {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	check := func(x1, y1, x2, y2 int) bool {
		return sum[x2][y2]-sum[x1-1][y2]-sum[x2][y1-1]+sum[x1-1][y1-1] <= threshold
	}

	r := min(m, n)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for c := ans + 1; c <= r; c++ {
				if i+c-1 <= m && j+c-1 <= n && check(i, j, i+c-1, j+c-1) {
					ans++
				} else {
					break
				}
			}
		}
	}

	return
}

// https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/solution/yuan-su-he-xiao-yu-deng-yu-yu-zhi-de-zheng-fang-2/
// @lc code=end
