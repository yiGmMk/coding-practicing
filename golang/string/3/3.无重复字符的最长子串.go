/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (38.93%)
 * Likes:    7969
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 4.9M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */
package jzoffer

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var res int
	cmap := make(map[rune]int, 26)

	left := 0
	for i, v := range s {
		if preId, ok := cmap[v]; ok {
			// 转移至重复字符的下一个位置,之前的标志全部删除
			for j := left; j < preId+1; j++ {
				delete(cmap, rune(s[j]))
			}
			left = preId + 1
		}
		// 更新最大值
		res = max(res, i-left+1)
		cmap[v] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
