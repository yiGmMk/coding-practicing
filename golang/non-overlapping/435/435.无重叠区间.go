/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 *
 * https://leetcode.cn/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (51.10%)
 * Likes:    810
 * Dislikes: 0
 * Total Accepted:    178.4K
 * Total Submissions: 349.1K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[1,3]]'
 *
 * 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回
 * 需要移除区间的最小数量，使剩余区间互不重叠 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
 * 输出: 1
 * 解释: 移除 [1,3] 后，剩下的区间没有重叠。
 *
 *
 * 示例 2:
 *
 *
 * 输入: intervals = [ [1,2], [1,2], [1,2] ]
 * 输出: 2
 * 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
 *
 *
 * 示例 3:
 *
 *
 * 输入: intervals = [ [1,2], [2,3] ]
 * 输出: 0
 * 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * -5 * 10^4 <= starti < endi <= 5 * 10^4
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start

type Pair struct {
	l, r int
}
type Pairs []Pair

func (ps Pairs) Less(i, j int) bool { return ps[i].l < ps[j].l }
func (ps Pairs) Len() int           { return len(ps) }
func (ps Pairs) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

// 按照右边界排序，就要从左向右遍历，因为右边界越小越好，只要右边界越小，留给下一个区间的空间就越大，所以从左向右遍历，优先选右边界小的
// 按照左边界排序，就要从右向左遍历，因为左边界数值越大越好（越靠右），这样就给前一个区间的空间就越大，所以可以从右向左遍历。
// 如果按照左边界排序，还从左向右遍历的话，要处理各个区间右边界的各种情况
// 这里没有直接去求移除区间的个数，而是转换一下思路，先去求非交叉区间的个数
// 用总个数 减去 非交叉区间个数 ，最终求得需要移除的区间个数
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	// 按右边界排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	// 不交叉区间个数
	ans, right := 1, intervals[0][1]

	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}

func eraseOverlapIntervalsTimeOut(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	// 按左边界排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// r <= l
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return n - max(dp...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/non-overlapping-intervals/solution/wu-zhong-die-qu-jian-by-leetcode-solutio-cpsb/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
