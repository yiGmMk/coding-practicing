/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.40%)
 * Likes:    854
 * Dislikes: 0
 * Total Accepted:    243K
 * Total Submissions: 547.2K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */
package jzoffer

// @lc code=start

// 每次都全量计算,有些重复计算了
func checkInclusion0(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	l := 0
	for r := l + len(s1); r <= len(s2); r++ {
		if check(s1, s2[l:r]) {
			return true
		}
		l++
	}
	return false
}

func check(s1 string, s2 string) bool {
	cnt := [26]int{}
	for i := 0; i < len(s1); i++ {
		cnt[s1[i]-'a']--
	}
	for i := 0; i < len(s2); i++ {
		cnt[s2[i]-'a']++
	}
	for _, v := range cnt {
		if v < 0 {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	cnt := [26]int{}
	for i := 0; i < len(s1); i++ {
		cnt[s1[i]-'a']--
	}
	check := func() bool {
		for _, v := range cnt {
			if v < 0 {
				return false
			}
		}
		return true
	}

	l := 0
	r := l + len(s1) // [l,r)
	for i := 0; i < r; i++ {
		cnt[s2[i]-'a']++
	}
	for ; r <= len(s2); r++ {
		if check() {
			return true
		}
		cnt[s2[l]-'a']--
		l++
		if r < len(s2) {
			cnt[s2[r]-'a']++
		}
	}
	return false
}

// @lc code=end
