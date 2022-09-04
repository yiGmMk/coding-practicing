/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode.cn/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (66.58%)
 * Likes:    952
 * Dislikes: 0
 * Total Accepted:    215.9K
 * Total Submissions: 324K
 * Testcase Example:  '"abc"'
 *
 * 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
 *
 * 回文字符串 是正着读和倒过来读一样的字符串。
 *
 * 子字符串 是字符串中的由连续字符组成的一个序列。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写英文字母组成
 *
 *
 */
package jzoffer

// @lc code=start
// O(n*n)*O(n)  遍历*判断
func countSubstringsN3(s string) int {
	var res int
	sm := make(map[string]bool, len(s))
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			sub := s[i:j]
			if v, ok := sm[sub]; ok {
				if v {
					res++
				}
				continue
			}
			v := is(sub)
			if v {
				res++
			}
			sm[sub] = v
		}
	}
	return res
}

func is(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

// ---------soluton from jzoffer-------------
// 遍历字符串，对每个字符，都看作回文的中心，向两端延申进行判断直到非回文。
// 回文的中心可能是一个字符，也可能是两个字符。
// 注意双指针可能越界。
func count(s string, st, ed int) int {
	var res int
	for st >= 0 && ed < len(s) && s[st] == s[ed] {
		res++
		st--
		ed++
	}
	return res
}

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += count(s, i, i)
		cnt += count(s, i, i+1)
	}
	return cnt
}

// @lc code=end
