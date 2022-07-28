/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode.cn/problems/add-binary/description/
 *
 * algorithms
 * Easy (53.78%)
 * Likes:    850
 * Dislikes: 0
 * Total Accepted:    262.4K
 * Total Submissions: 488K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串，返回它们的和（用二进制表示）。
 *
 * 输入为 非空 字符串且只包含数字 1 和 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 *
 *
 * 提示：
 *
 *
 * 每个字符串仅由字符 '0' 或 '1' 组成。
 * 1 <= a.length, b.length <= 10^4
 * 字符串如果不是 "0" ，就都不含前导零。
 *
 *
 */
package jzoffer

import "strconv"

// 11
//  1
//100
// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	res := make([]rune, 0, max(la, lb))
	var carry int
	getbit := func(ra, rb bool, carry int) (rune, int) {
		if ra && rb {
			if carry > 0 {
				return '1', carry
			}
			return '0', carry + 1
		}
		if ra || rb {
			if carry > 0 {
				return '0', carry
			}
			return '1', carry
		}
		if carry > 0 {
			return '1', carry - 1
		}
		return '0', carry
	}

	for i, j := la-1, lb-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var bit rune
		if i >= 0 && j >= 0 {
			bit, carry = getbit(rune(a[i]) == '1', rune(b[j]) == '1', carry)
			res = append(res, bit)
			continue
		}
		if i >= 0 {
			bit, carry = getbit(rune(a[i]) == '1', false, carry)
			res = append(res, bit)
			continue
		}
		if j >= 0 {
			bit, carry = getbit(false, rune(b[j]) == '1', carry)
			res = append(res, bit)
			continue
		}
	}
	if carry > 0 {
		res = append(res, '1')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

// @lc code=end

func addBinary2(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		ai, bi := lenA-i-1, lenB-i-1
		if i < lenA {
			carry += int(a[ai] - '0')
		}
		if i < lenB {
			carry += int(b[bi] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/JFETK5/solution/er-jin-zhi-jia-fa-by-leetcode-solution-fa6t/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
