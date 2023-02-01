/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode.cn/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (74.15%)
 * Likes:    513
 * Dislikes: 0
 * Total Accepted:    286.4K
 * Total Submissions: 386.2K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "Let's take LeetCode contest"
 * 输出："s'teL ekat edoCteeL tsetnoc"
 *
 *
 * 示例 2:
 *
 *
 * 输入： s = "God Ding"
 * 输出："doG gniD"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * s 包含可打印的 ASCII 字符。
 * s 不包含任何开头或结尾空格。
 * s 里 至少 有一个词。
 * s 中的所有单词都用一个空格隔开。
 *
 *
 */
package jzoffer

// @lc code=start
func reverseString(s []rune, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
func reverseWords(s string) string {
	l, r := -1, 0
	str := []rune(s)
	for r < len(s) {
		if str[r] == ' ' || r == len(str)-1 {
			if r != len(str)-1 && r-l > 1 {
				reverseString(str, l+1, r-1)
			} else {
				reverseString(str, l+1, r)
			}
			l = r
		}
		r++
	}
	return string(str)
}

// 找出一个词,翻转加入res,找空格,添加到res
func reverseWords1(s string) string {
	n := len(s)
	res := []byte{}
	for r := 0; r < n; {
		l := r
		for r < n && s[r] != ' ' {
			r++
		}
		for p := l; p < r; p++ {
			res = append(res, s[l+r-1-p]) // 翻转
		}
		for r < n && s[r] == ' ' {
			r++
			res = append(res, ' ')
		}
	}
	return string(res)
}

// @lc code=end
