/*
 * @lc app=leetcode.cn id=1657 lang=golang
 *
 * [1657] 确定两个字符串是否接近
 *
 * https://leetcode.cn/problems/determine-if-two-strings-are-close/description/
 *
 * algorithms
 * Medium (47.66%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    17.8K
 * Total Submissions: 37.3K
 * Testcase Example:  '"abc"\n"bca"'
 *
 * 如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：
 *
 *
 * 操作 1：交换任意两个 现有 字符。
 *
 *
 * 例如，abcde -> aecdb
 *
 *
 * 操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
 *
 * 例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
 *
 *
 *
 *
 * 你可以根据需要对任意一个字符串多次使用这两种操作。
 *
 * 给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "abc", word2 = "bca"
 * 输出：true
 * 解释：2 次操作从 word1 获得 word2 。
 * 执行操作 1："abc" -> "acb"
 * 执行操作 1："acb" -> "bca"
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "a", word2 = "aa"
 * 输出：false
 * 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
 *
 * 示例 3：
 *
 *
 * 输入：word1 = "cabbba", word2 = "abbccc"
 * 输出：true
 * 解释：3 次操作从 word1 获得 word2 。
 * 执行操作 1："cabbba" -> "caabbb"
 * 执行操作 2："caabbb" -> "baaccc"
 * 执行操作 2："baaccc" -> "abbccc"
 *
 *
 * 示例 4：
 *
 *
 * 输入：word1 = "cabbba", word2 = "aabbss"
 * 输出：false
 * 解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * word1 和 word2 仅包含小写英文字母
 *
 *
 */
package jzoffer

import (
	"reflect"
	"sort"
)

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	m1, m2 := make(map[rune]int), make(map[rune]int)
	for _, v := range word1 {
		m1[v]++
	}
	for _, v := range word2 {
		m2[v]++
	}
	// 种类数量相同
	if len(m1) != len(m2) {
		return false
	}
	ca, cb := make([]int, len(m1)), make([]int, len(m1))
	idx := 0
	// 频率相同,两个字符串都有同样的类型
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
		ca[idx] = v
		idx++
	}
	idx = 0
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			return false
		}
		cb[idx] = v
		idx++
	}
	sort.Ints(ca)
	sort.Ints(cb)
	if !reflect.DeepEqual(ca, cb) {
		return false
	}
	return true
}

// @lc code=end
