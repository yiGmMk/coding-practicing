/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (57.71%)
 * Likes:    3334
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.2M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 *
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 *
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[7,1,5,3,6,4]
 * 输出：5
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */
package jzoffer

import "math"

// @lc code=start
// 只能买卖一次,一次能获取最大利润的卖出一定是在卖出前(买到买入最低价),维护卖出前的最低价即可
func maxProfitGreedy(prices []int) (res int) {
	if len(prices) < 2 {
		return 0
	}
	minPrice := math.MaxInt
	for _, v := range prices {
		minPrice = min(minPrice, v)
		res = max(res, v-minPrice)
	}
	return
}

// dp解法 dp[i][j]表示下标为 i 这一天结束的时候，手上持股状态为 j 时，我们持有的现金数
// j = 0，表示当前不持股；
// j = 1，表示当前持股。
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 作者：liweiwei1419
// 链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func maxProfit(prices []int) (res int) {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// dp[i][0] 下标为 i 这天结束的时候，不持股，手上拥有的现金数
	// dp[i][1] 下标为 i 这天结束的时候，持股，手上拥有的现金数

	// 初始化：不持股显然为 0，持股就需要减去第 1 天（下标为 0）的股价
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]

	// 如何保证只交易一次呢?
	// dp[i][1] = max(dp[i - 1][1], -prices[i]); 今天的持股状态只能从昨天的持股状态
	// 或者没有买入前的不持股状态转化而来的，可见
	// 今天的持股状态dp[i][1]不能从昨天的不持股状态dp[i-1][0]转化而来，
	// 即昨天如果不持股，今天也不会持股，保证了买入只有一次
}

// @lc code=end
