/*
 * @lc app=leetcode.cn id=2180 lang=golang
 *
 * [2180] 统计各位数字之和为偶数的整数个数
 *
 * https://leetcode.cn/problems/count-integers-with-even-digit-sum/description/
 *
 * algorithms
 * Easy (64.86%)
 * Likes:    47
 * Dislikes: 0
 * Total Accepted:    23.7K
 * Total Submissions: 34.6K
 * Testcase Example:  '4'
 *
 * 给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。
 *
 * 正整数的 各位数字之和 是其所有位上的对应数字相加的结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 4
 * 输出：2
 * 解释：
 * 只有 2 和 4 满足小于等于 4 且各位数字之和为偶数。
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = 30
 * 输出：14
 * 解释：
 * 只有 14 个整数满足小于等于 30 且各位数字之和为偶数，分别是：
 * 2、4、6、8、11、13、15、17、19、20、22、24、26 和 28 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 1000
 *
 *
 */
package jzoffer

// @lc code=start
func countEven1(num int) int {
	check := func(x int) bool {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		return sum%2 == 0
	}

	res := 0
	for i := 2; i <= num; i++ {
		if check(i) {
			res++
		}
	}
	return res
}

// 【C++】数学
// 不妨大胆猜测，奇偶是均匀分布的。仔细验证一下，每十个数（xx0 ~ xx9）必然5奇5偶并且交替分布。
// 因此最后奇偶数量必然是相同或者差1。那么只需要分析是谁多即可
// 可以看出，如果n是奇数，那么必然是奇数比偶数多一个。否则取决于n的数位和的奇偶性。
// https://leetcode.cn/problems/count-integers-with-even-digit-sum/solution/tong-ji-ge-wei-shu-zi-zhi-he-wei-ou-shu-rvqol/
// 1 2 3 4 5 6 7 8 9 10
func countEven(num int) int {
	if num&1 > 0 { // num为奇数,有num/2个
		return num / 2
	}
	v, sum := num, 0
	for v > 0 {
		sum += v % 10
		v /= 10
	}
	return num/2 - sum&1
}

// @lc code=end
