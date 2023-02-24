/*
 * @lc app=leetcode.cn id=1238 lang=golang
 *
 * [1238] 循环码排列
 *
 * https://leetcode.cn/problems/circular-permutation-in-binary-representation/description/
 *
 * algorithms
 * Medium (67.63%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    8.7K
 * Total Submissions: 11.6K
 * Testcase Example:  '2\n3'
 *
 * 给你两个整数 n 和 start。你的任务是返回任意 (0,1,2,,...,2^n-1) 的排列 p，并且满足：
 *
 *
 * p[0] = start
 * p[i] 和 p[i+1] 的二进制表示形式只有一位不同
 * p[0] 和 p[2^n -1] 的二进制表示形式也只有一位不同
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 2, start = 3
 * 输出：[3,2,0,1]
 * 解释：这个排列的二进制表示是 (11,10,00,01)
 * ⁠    所有的相邻元素都有一位是不同的，另一个有效的排列是 [3,1,0,2]
 *
 *
 * 示例 2：
 *
 * 输出：n = 3, start = 2
 * 输出：[2,6,7,5,4,0,1,3]
 * 解释：这个排列的二进制表示是 (010,110,111,101,100,000,001,011)
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 16
 * 0 <= start < 2^n
 *
 *
 */
package jzoffer

// @lc code=start
func check(s1, s2 string) bool {
	has := false
	for i, v := range s1 {
		if byte(v) == s2[i] {
			continue
		}
		if has {
			return false
		} else {
			has = true
		}
	}
	return true
}

/*
	格雷码(gray code)

格雷码是一种循环二进制码或者叫作反射二进制码。
格雷码的特点是从一个数变为相邻的一个数时，只有一个数据位发生跳变，
由于这种特点，就可以避免二进制编码计数组合电路中出现的亚稳态。
格雷码常用于通信，FIFO 或者 RAM 地址寻址计数器中。
格雷码属于可靠性编码，是一种错误最小化的编码方式，因为虽然二进制码可以直接由数/模转换器转换成模拟信号，
但在某些情况，例如从十进制的 3 转换为 4 时二进制码的每一位都要变，能使数字电路产生很大的尖峰电流脉冲。
而格雷码则没有这一缺点，它在相邻位间转换时，只有一位产生变化。它大大地减少了由一个状态到下一个状态时逻辑的混淆
*/
func genGrayCode(n int) []int {
	num := 1 << n
	res := make([]int, num)

	for i := 1; i < num; i++ {
		res[i] = i>>1 ^ i
	}
	return res
}

// 与Leetcode 89.格雷编码类似
func circularPermutation1(n int, start int) []int {
	res := genGrayCode(n)
	i := 0
	for ; i < len(res); i++ {
		if res[i] == start {
			break
		}
	}

	return append(res[i:], res[:i]...)
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/circular-permutation-in-binary-representation/solution/xun-huan-ma-pai-lie-by-leetcode-solution-6e40/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func circularPermutation(n int, start int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = (i >> 1) ^ i ^ start
	}
	return ans
}

// @lc code=end
