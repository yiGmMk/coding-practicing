/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 *
 * https://leetcode.cn/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (73.66%)
 * Likes:    368
 * Dislikes: 0
 * Total Accepted:    59.3K
 * Total Submissions: 80.5K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * 给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j])
 * 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
 * 输出：16
 * 解释：这两个单词为 "abcw", "xtfn"。
 *
 * 示例 2：
 *
 *
 * 输入：words = ["a","ab","abc","d","cd","bcd","abcd"]
 * 输出：4
 * 解释：这两个单词为 "ab", "cd"。
 *
 * 示例 3：
 *
 *
 * 输入：words = ["a","aa","aaa","aaaa"]
 * 输出：0
 * 解释：不存在这样的两个单词。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= words.length <= 1000
 * 1 <= words[i].length <= 1000
 * words[i] 仅包含小写字母
 *
 *
 */
package jzoffer

import "github.com/yiGmMk/leetcode/golang/util"

//只有26个字母,字符串中出现的字母可以使用int标识32
// @lc code=start
func maxProduct01(words []string) int {
	var res int32
	cm := make(map[string]int, len(words))
	for _, word := range words {
		mask := 0 // 64位的掩码
		for _, c := range word {
			mask |= 1 << uint(c-'a')
		}
		cm[word] = mask
		for k, v := range cm {
			if (v & mask) == 0 {
				res = util.Max(res, int32(len(word)*len(k)))
			}
		}
	}
	return int(res)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// 优化内存占用
func maxProduct(words []string) int {
	var res int32
	cm := make(map[int]int, len(words)) // mask:len(word)
	for _, word := range words {
		mask := 0 // 64位的掩码
		for _, c := range word {
			mask |= 1 << uint(c-'a')
		}
		// 取最大长度,mask一致的单词可能有多个,长度不一致
		if len(word) > cm[mask] {
			cm[mask] = len(word)
		}

		for k, v := range cm {
			if (k & mask) == 0 {
				res = max(res, int32(len(word)*v))
			}
		}
	}
	return int(res)
}

// @lc code=end
