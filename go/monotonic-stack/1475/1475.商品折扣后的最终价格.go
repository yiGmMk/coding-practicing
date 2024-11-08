/*
 * @lc app=leetcode.cn id=1475 lang=golang
 *
 * [1475] 商品折扣后的最终价格
 *
 * https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/description/
 *
 * algorithms
 * Easy (73.31%)
 * Likes:    207
 * Dislikes: 0
 * Total Accepted:    70.4K
 * Total Submissions: 95.9K
 * Testcase Example:  '[8,4,6,2,3]'
 *
 * 给你一个数组 prices ，其中 prices[i] 是商店里第 i 件商品的价格。
 *
 * 商店里正在进行促销活动，如果你要买第 i 件商品，那么你可以得到与 prices[j] 相等的折扣，其中 j 是满足 j > i 且 prices[j]
 * <= prices[i] 的 最小下标 ，如果没有满足条件的 j ，你将没有任何折扣。
 *
 * 请你返回一个数组，数组中第 i 个元素是折扣后你购买商品 i 最终需要支付的价格。
 *
 *
 *
 * 示例 1：
 *
 * 输入：prices = [8,4,6,2,3]
 * 输出：[4,2,4,2,3]
 * 解释：
 * 商品 0 的价格为 price[0]=8 ，你将得到 prices[1]=4 的折扣，所以最终价格为 8 - 4 = 4 。
 * 商品 1 的价格为 price[1]=4 ，你将得到 prices[3]=2 的折扣，所以最终价格为 4 - 2 = 2 。
 * 商品 2 的价格为 price[2]=6 ，你将得到 prices[3]=2 的折扣，所以最终价格为 6 - 2 = 4 。
 * 商品 3 和 4 都没有折扣。
 *
 *
 * 示例 2：
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：[1,2,3,4,5]
 * 解释：在这个例子中，所有商品都没有折扣。
 *
 *
 * 示例 3：
 *
 * 输入：prices = [10,1,1,6]
 * 输出：[9,0,1,6]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 500
 * 1 <= prices[i] <= 10^3
 *
 *
 */
package jzoffer

// @lc code=start

// 单调栈 找后面比当前小的最近元素
// 正序遍历,单调栈维护还没有找到折扣价的
func finalPrices(prices []int) []int {
	var (
		idx   = -1
		stack = []struct{ i, v int }{}
	)
	for _, v := range prices {
		idx++
		for len(stack) > 0 && v <= stack[0].v {
			low := stack[0]
			prices[low.i] = low.v - v
			stack = stack[1:]
		}
		stack = append([]struct {
			i int
			v int
		}{{i: idx, v: v}}, stack...)
	}
	return prices
}

// 倒序遍历,单调栈:维护比当前位置小的右侧元素
// func finalPrices(prices []int) []int {
//     n := len(prices)
//     ans := make([]int, n)
//     st := []int{0}
//     for i := n - 1; i >= 0; i-- {
//         p := prices[i]
//         for len(st) > 1 && st[len(st)-1] > p {
//             st = st[:len(st)-1]
//         }
//         ans[i] = p - st[len(st)-1]
//         st = append(st, p)
//     }
//     return ans
// }
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
