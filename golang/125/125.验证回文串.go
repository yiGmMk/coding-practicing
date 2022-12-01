/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.90%)
 * Likes:    557
 * Dislikes: 0
 * Total Accepted:    391K
 * Total Submissions: 833.9K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 解释："amanaplanacanalpanama" 是回文串
 *
 *
 * 示例 2:
 *
 *
 * 输入: "race a car"
 * 输出: false
 * 解释："raceacar" 不是回文串
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 字符串 s 由 ASCII 字符组成
 *
 *
 */
package jzoffer

import (
	"strings"
	"unicode"
)

// @lc code=start
func isPalindrome3(s string) bool {
	if s == "" {
		return true
	}
	cnt := 0
	trim := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}
		cnt++
		return -1
	}, s)
	// 反转与原始相同

	reverse := []rune(trim)
	for i, j := 0, len(reverse)-1; i < j; {
		reverse[i], reverse[j] = reverse[j], reverse[i]
		i++
		j--
	}

	// log.Printf("%s,%s\n", trim, string(reverse))
	return trim == string(reverse)
}

func isPalindrome1(s string) bool {
	if s == "" {
		return true
	}
	cnt := 0
	trim := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}
		cnt++
		return -1
	}, s)
	// 反转与原始相同
	//log.Println(len(s), cnt, len(trim), trim)

	reverse := []rune{}
	for i := len(trim) - 1; i >= 0; i-- {
		reverse = append(reverse, rune(trim[i]))
	}

	//log.Printf("%s,%s\n", trim, string(reverse))
	return trim == string(reverse)
}

func isPalindrome2(s string) bool {
	if s == "" {
		return true
	}
	for i, j := 0, len(s)-1; i < j; {
		for !(unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))) && i < j {
			i++
		}
		for !(unicode.IsLetter(rune(s[j])) || unicode.IsNumber(rune(s[j]))) && j > 0 {
			j--
		}
		if i >= j {
			break
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		} else {
			i++
			j--
		}
	}

	return true
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	valid := func(r rune) bool {
		return (unicode.IsLetter(r) || unicode.IsNumber(r))
	}
	for i, j := 0, len(s)-1; i < j; {
		si, sj := unicode.ToLower(rune(s[i])), unicode.ToLower(rune(s[j]))
		for !valid(si) && i < j {
			i++
			si = unicode.ToLower(rune(s[i]))
		}
		for !valid(sj) && j > 0 {
			j--
			sj = unicode.ToLower(rune(s[j]))
		}
		if i >= j {
			break
		}
		if unicode.ToLower(si) != unicode.ToLower(sj) {
			return false
		} else {
			i++
			j--
		}
	}

	return true
}

// @lc code=end
