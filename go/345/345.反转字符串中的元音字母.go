/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode.cn/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (54.53%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    177.4K
 * Total Submissions: 325.4K
 * Testcase Example:  '"hello"'
 *
 * 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
 *
 * 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "hello"
 * 输出："holle"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "leetcode"
 * 输出："leotcede"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 3 * 10^5
 * s 由 可打印的 ASCII 字符组成
 *
 *
 */
package jzoffer

// @lc code=start
func check(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
func reverseVowels(s string) string {
	res := []byte(s)
	for l, r := 0, len(s)-1; l <= r; {
		for ; l <= r && !check(res[l]); l++ {
		}
		for ; r >= l && !check(res[r]); r-- {
		}
		if l <= r && check(res[l]) && check(res[r]) {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
	}

	return string(res)
}

// @lc code=end
