/*
 * @lc app=leetcode.cn id=1250 lang=golang
 *
 * [1250] 检查「好数组」
 *
 * https://leetcode.cn/problems/check-if-it-is-a-good-array/description/
 *
 * algorithms
 * Hard (59.71%)
 * Likes:    65
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 14.2K
 * Testcase Example:  '[12,5,7,23]'
 *
 * 给你一个正整数数组 nums，你需要从中任选一些子集，然后将子集中每一个数乘以一个 任意整数，并求出他们的和。
 *
 * 假如该和结果为 1，那么原数组就是一个「好数组」，则返回 True；否则请返回 False。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [12,5,7,23]
 * 输出：true
 * 解释：挑选数字 5 和 7。
 * 5*3 + 7*(-2) = 1
 *
 *
 * 示例 2：
 *
 * 输入：nums = [29,6,10]
 * 输出：true
 * 解释：挑选数字 29, 6 和 10。
 * 29*1 + 6*(-3) + 10*(-1) = 1
 *
 *
 * 示例 3：
 *
 * 输入：nums = [3,6]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 *
 *
 */
package jzoffer

// @lc code=start

// 1. Eq. ax+by=1 has solution x, y if gcd(a,b) = 1.
// 看数组中是否有部分子数组的最大公约数==1
// 求nums 中全部数字的最大公约数的方法为，我们设初始为
// x=gcd(x,nums[i])。遍历完全部数字后，x 即为数组
// nums 中全部的元素的最大公约数。然后判断其是否等于
// 1 即可。在实现过程中我们也可以进一步做优化：如果遍历过程中出现最大公约数等于
// 1 的情况，则由于 1 和任何正整数的最大公约数都是
// 1，此时可以提前结束遍历。
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/check-if-it-is-a-good-array/solution/jian-cha-hao-shu-zu-by-leetcode-solution-qg2h/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func isGoodArray(nums []int) bool {
	g := 0
	for _, v := range nums {
		g = gcd(g, v)
		if g == 1 {
			return true
		}
	}
	return false
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end
