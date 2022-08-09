/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (44.58%)
 * Likes:    2052
 * Dislikes: 0
 * Total Accepted:    326.8K
 * Total Submissions: 732.1K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
 */
package jzoffer

import "math"

// @lc code=start
// 超时
func minWindowTimeout(s string, t string) string {
	var res string
	cmap := make(map[rune]int, len(t)) // 实际需要比较的数据远小于len(t),这就是超时的原因
	for _, c := range t {
		cmap[c]++
	}

	check := func(cur, tar map[rune]int) bool {
		for k, v := range tar {
			if cur[k] < v {
				return false
			}
		}
		return true
	}

	left := 0
	num := len(t)
	has, curNum := make(map[rune]int, len(t)), 0 // 实际需要比较的数据远小于len(t)
	for right, v := range s {
		if cmap[v] <= 0 {
			curNum++
			continue
		}
		has[v]++
		curNum++
		if curNum >= num {
			for check(has, cmap) && left <= right {
				if res != "" {
					res = min(res, s[left:right+1])
				} else {
					res = s[left : right+1]
				}

				has[rune(s[left])]--
				left++
				curNum--
			}
		}
	}

	return res
}

func minWindowFixTimeout(s string, t string) string {
	var res string
	cmap := make(map[rune]int /*, len(t)*/) // 实际需要比较的数据远小于len(t),这就是超时的原因
	for _, c := range t {
		cmap[c]++
	}

	check := func(cur, tar map[rune]int) bool {
		for k, v := range tar {
			if cur[k] < v {
				return false
			}
		}
		return true
	}

	left := 0
	num := len(t)
	has, curNum := make(map[rune]int /*, len(t)*/), 0 // 实际需要比较的数据远小于len(t)
	for right, v := range s {
		if cmap[v] <= 0 {
			curNum++
			continue
		}
		has[v]++
		curNum++
		if curNum >= num {
			for check(has, cmap) && left <= right {
				if res != "" {
					res = min(res, s[left:right+1])
				} else {
					res = s[left : right+1]
				}

				has[rune(s[left])]--
				left++
				curNum--
			}
		}
	}

	return res
}

func min(a, b string) string {
	n, m := len(a), len(b)
	if n < m {
		return a
	}
	return b
}

func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)
	for _, c := range t {
		tMap[byte(c)]++
	}

	check := func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	ansL, ansR, ansLen := -1, -1, math.MaxInt32

	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && tMap[s[r]] > 0 {
			sMap[s[r]]++
		}

		for check() && l <= r {
			if r-l+1 < ansLen {
				ansLen = r - l + 1
				ansL, ansR = l, l+ansLen
			}

			if _, ok := tMap[s[l]]; ok {
				sMap[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}

// @lc code=end
