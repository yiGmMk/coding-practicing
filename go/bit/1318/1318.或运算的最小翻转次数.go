/*
 * @lc app=leetcode.cn id=1318 lang=golang
 *
 * [1318] 或运算的最小翻转次数
 *
 * https://leetcode.cn/problems/minimum-flips-to-make-a-or-b-equal-to-c/description/
 *
 * algorithms
 * Medium (67.45%)
 * Likes:    61
 * Dislikes: 0
 * Total Accepted:    12.9K
 * Total Submissions: 19.1K
 * Testcase Example:  '2\n6\n5'
 *
 * 给你三个正整数 a、b 和 c。
 *
 * 你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。
 *
 * 「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：a = 2, b = 6, c = 5
 * 输出：3
 * 解释：翻转后 a = 1 , b = 4 , c = 5 使得 a OR b == c
 *
 * 示例 2：
 *
 * 输入：a = 4, b = 2, c = 7
 * 输出：1
 *
 *
 * 示例 3：
 *
 * 输入：a = 1, b = 2, c = 3
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= a <= 10^9
 * 1 <= b <= 10^9
 * 1 <= c <= 10^9
 *
 *
 */

package jzoffer

import "fmt"

// @lc code=start
func minFlipsUsingString(a int, b int, c int) (res int) {
	trans := func(v int) string {
		return fmt.Sprintf("%032b", v)
	}
	sa, sb, sc := trans(a), trans(b), trans(c)
	for i := 0; i < len(sa); i++ {
		if sc[i] == '0' {
			if sa[i] == '1' {
				res++
			}
			if sb[i] == '1' {
				res++
			}
			continue
		}
		if sa[i] == '0' && sb[i] == '0' {
			res++
		}
	}
	return
}

// 不用string, 位运算
func minFlips1(a int, b int, c int) (res int) {
	for i := 0; i < 32; i++ {
		// 1. c的第i位为0
		if c&(1<<i) == 0 {
			// a的第i位为1
			if a&(1<<i) != 0 {
				res++
			}
			// b的第i位为1
			if b&(1<<i) != 0 {
				res++
			}
			continue
		}
		// 2. c的第i位为1
		if a&(1<<i) == 0 && b&(1<<i) == 0 {
			res++
		}
	}
	return
}

// @lc code=end
