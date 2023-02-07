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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring0(s string) int {
	l, r, res, vis := 0, 0, 0, make(map[byte]int)
	for r < len(s) {
		if i, ok := vis[s[r]]; ok {
			res = max(res, r-l)
			for id := l; id < i+1; id++ { //不能从零开始删除,删多了
				delete(vis, s[id])
			}
			l = i + 1
		} else {
			res = max(res, r-l+1)
		}
		vis[s[r]] = r
		r++
	}
	if len(vis) == len(s) { // 没有重复的
		return len(s)
	}
	return res
}

func lengthOfLongestSubstring1(s string) int {
	l, res, vis := 0, 0, make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if i, ok := vis[s[r]]; ok {
			res = max(res, r-l)
			for id := l; id < i+1; id++ { //不能从零开始删除,删多了
				delete(vis, s[id])
			}
			l = i + 1
		} else {
			res = max(res, r-l+1)
		}
		vis[s[r]] = r
	}
	if len(vis) == len(s) { // 没有重复的
		return len(s)
	}
	return res
}

// TODO
// 什么是滑动窗口？
// 其实就是一个队列,比如例题中的 abcabcbb，进入这个队列（窗口）为 abc 满足题目要求，当再进入 a，队列变成了 abca，这时候不满足要求。所以，我们要移动这个队列！
// 如何移动？
// 我们只要把队列的左边的元素移出就行了，直到满足题目要求！
// 一直维持这样的队列，找出队列出现最长的长度时候，求出解！
// 时间复杂度：O(n)
// 作者：powcai
// 链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func lengthOfLongestSubstring(s string) int {
	l, res, vis := 0, 0, make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if i, ok := vis[s[r]]; ok {
			l = max(l, i+1)
		}
		res = max(res, r-l+1)
		vis[s[r]] = r
	}
	if len(vis) == len(s) { // 没有重复的
		return len(s)
	}
	return res
}

// @lc code=end
