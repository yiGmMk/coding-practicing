/*
 * @lc app=leetcode.cn id=2517 lang=golang
 *
 * [2517] 礼盒的最大甜蜜度
 *
 * https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/description/
 *
 * algorithms
 * Medium (70.12%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 5.4K
 * Testcase Example:  '[13,5,1,8,21,2]\n3'
 *
 * 给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。
 *
 * 商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。
 *
 * 返回礼盒的 最大 甜蜜度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：price = [13,5,1,8,21,2], k = 3
 * 输出：8
 * 解释：选出价格分别为 [13,5,21] 的三类糖果。
 * 礼盒的甜蜜度为 min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8 。
 * 可以证明能够取得的最大甜蜜度就是 8 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：price = [1,3,1], k = 2
 * 输出：2
 * 解释：选出价格分别为 [1,3] 的两类糖果。
 * 礼盒的甜蜜度为 min(|1 - 3|) = min(2) = 2 。
 * 可以证明能够取得的最大甜蜜度就是 2 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：price = [7,7,7,7], k = 2
 * 输出：0
 * 解释：从现有的糖果中任选两类糖果，甜蜜度都会是 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= price.length <= 10^5
 * 1 <= price[i] <= 10^9
 * 2 <= k <= price.length
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/solution/er-fen-da-an-by-endlesscheng-r418/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)

	// 从小(x)往大,下一个可选的数是x+d,算可选的数cnt,cnt<k说明d取大了
	// 题目要求取k类不同的
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		d++
		cnt, x0 := 1, price[0]
		for _, x := range price[1:] {
			if x >= x0+d {
				cnt++
				x0 = x
			}
		}
		return cnt < k
	})
}

// @lc code=end
