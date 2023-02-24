/*
 * @lc app=leetcode.cn id=1326 lang=golang
 *
 * [1326] 灌溉花园的最少水龙头数目
 *
 * https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/description/
 *
 * algorithms
 * Hard (48.39%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    17.2K
 * Total Submissions: 32.5K
 * Testcase Example:  '5\n[3,4,1,1,0,0]'
 *
 * 在 x 轴上有一个一维的花园。花园长度为 n，从点 0 开始，到点 n 结束。
 *
 * 花园里总共有 n + 1 个水龙头，分别位于 [0, 1, ..., n] 。
 *
 * 给你一个整数 n 和一个长度为 n + 1 的整数数组 ranges ，其中 ranges[i] （下标从 0 开始）表示：如果打开点 i
 * 处的水龙头，可以灌溉的区域为 [i -  ranges[i], i + ranges[i]] 。
 *
 * 请你返回可以灌溉整个花园的 最少水龙头数目 。如果花园始终存在无法灌溉到的地方，请你返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：n = 5, ranges = [3,4,1,1,0,0]
 * 输出：1
 * 解释：
 * 点 0 处的水龙头可以灌溉区间 [-3,3]
 * 点 1 处的水龙头可以灌溉区间 [-3,5]
 * 点 2 处的水龙头可以灌溉区间 [1,3]
 * 点 3 处的水龙头可以灌溉区间 [2,4]
 * 点 4 处的水龙头可以灌溉区间 [4,4]
 * 点 5 处的水龙头可以灌溉区间 [5,5]
 * 只需要打开点 1 处的水龙头即可灌溉整个花园 [0,5] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 3, ranges = [0,0,0,0]
 * 输出：-1
 * 解释：即使打开所有水龙头，你也无法灌溉整个花园。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 * ranges.length == n + 1
 * 0 <= ranges[i] <= 100
 *
 *
 */

package jzoffer

import (
	"math"
	"sort"
)

// @lc code=start

/*
 1. Create intervals of the area covered by each tap, sort intervals by the left end
 2. We need to cover the interval [0, n]. we can start with the first interval and
    out of all intervals that intersect with it we choose the one that covers the
    farthest point to the right.
 3. What if there is a gap between intervals that is not covered ? we should stop
    and return -1 as there is some interval that cannot be covered.
*/
func minTapsGreedy(n int, ranges []int) int {
	// 每次贪心找右边最远的
	// https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/comments/392185
	intervals := [][]int{}
	for i, v := range ranges {
		intervals = append(intervals, []int{i - v, i + v})
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	i, res, r := 0, 0, 0 // r:当前覆盖范围的右端点
	for r < n {
		maxR := -1 // 下一个可覆盖的右端点
		for i < len(ranges) && intervals[i][0] <= r {
			maxR = max(maxR, intervals[i][1])
			i++
		}
		if maxR == -1 {
			return -1
		}
		r = maxR
		res++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i]为满足区间覆盖需要的最小数目
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/solution/guan-gai-hua-yuan-de-zui-shao-shui-long-tou-shu-3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func minTaps(n int, ranges []int) int {
	intervals := [][2]int{}
	for i, r := range ranges {
		start := max(0, i-r)
		end := min(n, i+r)
		intervals = append(intervals, [2]int{start, end})
	}
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0]
	})

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for _, p := range intervals {
		start, end := p[0], p[1]
		if dp[start] == math.MaxInt {
			return -1
		}
		for j := start; j <= end; j++ {
			dp[j] = min(dp[j], dp[start]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/solution/yi-zhang-tu-miao-dong-pythonjavacgo-by-e-wqry/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func minTapsDP(n int, ranges []int) (ans int) {
	rightMost := make([]int, n+1)
	for i, r := range ranges {
		left := max(i-r, 0)
		rightMost[left] = max(rightMost[left], i+r)
	}

	curRight := 0                     // 已建造的桥的右端点
	nextRight := 0                    // 下一座桥的右端点的最大值
	for i, r := range rightMost[:n] { // 注意这里没有遍历到 n，因为它已经是终点了
		nextRight = max(nextRight, r)
		if i == curRight { // 到达已建造的桥的右端点
			if i == nextRight { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curRight = nextRight // 造一座桥
			ans++
		}
	}
	return
}

// TODO:
// 相似题目
// 55. 跳跃游戏
// 45. 跳跃游戏 II
// 1024. 视频拼接
// @lc code=end
