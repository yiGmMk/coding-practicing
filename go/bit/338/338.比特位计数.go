/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 *
 * https://leetcode.cn/problems/counting-bits/description/
 *
 * algorithms
 * Easy (78.65%)
 * Likes:    1280
 * Dislikes: 0
 * Total Accepted:    316.7K
 * Total Submissions: 402.6K
 * Testcase Example:  '2'
 *
 * 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans
 * 作为答案。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：[0,1,1]
 * 解释：
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 5
 * 输出：[0,1,1,2,1,2]
 * 解释：
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 * 3 --> 11
 * 4 --> 100
 * 5 --> 101
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 10^5
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
 * 你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）
 *
 *
 *
 *
 */
package jzoffer

import (
	"math/bits"
)

// @lc code=start
func countBitsBuiltIn(n int) (res []int) {

	// fmt.Printf("%b,%b,%b", 63&(1<<7-1), (1<<7 - 1), 1<<7)

	res = make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = bits.OnesCount32(uint32(i))
	}
	return res
}

// 使用Brain Kcrnighan算法每次将最低位的二进制位1置为0，直到x为0
// 如 二进制数 101 &  100 => 100       11  &  10 => 10
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBitsBKAlgorithm(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

// 动态规划,上面的BK算法需要O(nlog(n))时间复杂度
// 最低设置位 x & (x-1) 会将最低位的1置为0
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

// @lc code=end
