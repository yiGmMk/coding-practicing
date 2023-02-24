/*
 * @lc app=leetcode.cn id=1124 lang=golang
 *
 * [1124] 表现良好的最长时间段
 *
 * https://leetcode.cn/problems/longest-well-performing-interval/description/
 *
 * algorithms
 * Medium (34.80%)
 * Likes:    442
 * Dislikes: 0
 * Total Accepted:    38.6K
 * Total Submissions: 99.5K
 * Testcase Example:  '[9,9,6,0,6,6,9]'
 *
 * 给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
 *
 * 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
 *
 * 所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
 *
 * 请你返回「表现良好时间段」的最大长度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：hours = [9,9,6,0,6,6,9]
 * 输出：3
 * 解释：最长的表现良好时间段是 [9,9,6]。
 *
 * 示例 2：
 *
 *
 * 输入：hours = [6,6,6]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= hours.length <= 10^4
 * 0 <= hours[i] <= 16
 *
 *
 */
package jzoffer

// @lc code=start

// TODO: 前缀和+单调栈
// 大于8的为1,小于8的为-1,先求前缀和将问题转化成求前缀和单调递增的最长范围.
func longestWPI(hours []int) int {
	for i, v := range hours {
		if v > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	prefixSum := make([]int, len(hours)+1)
	mStack := []int{0}
	for i := 1; i <= len(hours); i++ {
		prefixSum[i] = prefixSum[i-1] + hours[i-1]
		if prefixSum[i] < prefixSum[mStack[len(mStack)-1]] {
			mStack = append(mStack, i)
		}
	}

	// ms中的都是拐点
	// 从后往前遍历,求最长递增
	res := 0
	for i := len(hours); i > 0; i-- {
		for len(mStack) > 0 && prefixSum[i] > prefixSum[mStack[len(mStack)-1]] {
			res = max(res, i-mStack[len(mStack)-1]) // [栈顶,i) 可能是最长子数组
			mStack = mStack[:len(mStack)-1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/longest-well-performing-interval/solution/liang-chong-zuo-fa-liang-zhang-tu-miao-d-hysl/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func longestWPI1(hours []int) (ans int) {
	n := len(hours)
	s := make([]int, n+1) // 前缀和
	st := []int{0}        // s[0]
	for j, h := range hours {
		j++
		s[j] = s[j-1]
		if h > 8 {
			s[j]++
		} else {
			s[j]--
		}
		if s[j] < s[st[len(st)-1]] {
			st = append(st, j) // 感兴趣的 j
		}
	}
	for i := n; i > 0; i-- {
		for len(st) > 0 && s[i] > s[st[len(st)-1]] {
			ans = max(ans, i-st[len(st)-1]) // [栈顶,i) 可能是最长子数组
			st = st[:len(st)-1]
		}
	}
	return
}

// @lc code=end
