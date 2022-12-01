/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 *
 * https://leetcode.cn/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (40.10%)
 * Likes:    524
 * Dislikes: 0
 * Total Accepted:    116.9K
 * Total Submissions: 291.7K
 * Testcase Example:  '"aba"'
 *
 * 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "aba"
 * 输出: true
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "abca"
 * 输出: true
 * 解释: 你可以删除c字符。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "abc"
 * 输出: false
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * s 由小写英文字母组成
 *
 *
 */
package jzoffer

// @lc code=start

// 465/469,部分通不过
func validPalindromeBug(s string) bool {
	equal := func(i, j int) bool {
		if i >= len(s) || j < 0 {
			return false
		}
		return s[i] == s[j]
	}
	ss := []rune(s)
	remove := 0
	for i, j := 0, len(ss)-1; i < j; {
		if remove > 1 {
			return false
		}
		if !equal(i, j) {
			if equal(i+1, j) {
				remove++
				i++
				continue
			} else if equal(i, j-1) {
				remove++
				j--
				continue
			} else {
				return false
			}
		}
		if remove > 1 {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome1(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

func validPalindrome(s string) bool {
	equal := func(i, j int) bool {
		if i >= len(s) || j < 0 {
			return false
		}
		return s[i] == s[j]
	}
	for i, j := 0, len(s)-1; i < j; {
		if equal(i, j) {
			i++
			j--
			continue
		}
		i1, j1, flag1 := i+1, j, true
		for i1 < j1 {
			if !equal(i1, j1) {
				flag1 = false
				break
			}
			i1++
			j1--
		}
		i2, j2, flag2 := i, j-1, true
		for i2 < j2 {
			if !equal(i2, j2) {
				flag2 = false
				break
			}
			i2++
			j2--
		}
		return flag1 || flag2
	}
	return true
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/RQku0D/solution/zui-duo-shan-chu-yi-ge-zi-fu-de-dao-hui-30b55/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
