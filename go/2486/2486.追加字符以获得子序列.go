/*
 * @lc app=leetcode.cn id=2486 lang=golang
 *
 * [2486] 追加字符以获得子序列
 *
 * https://leetcode.cn/problems/append-characters-to-string-to-make-subsequence/description/
 *
 * algorithms
 * Medium (64.36%)
 * Likes:    7
 * Dislikes: 0
 * Total Accepted:    6.2K
 * Total Submissions: 9.6K
 * Testcase Example:  '"coaching"\n"coding"'
 *
 * 给你两个仅由小写英文字母组成的字符串 s 和 t 。
 *
 * 现在需要通过向 s 末尾追加字符的方式使 t 变成 s 的一个 子序列 ，返回需要追加的最少字符数。
 *
 * 子序列是一个可以由其他字符串删除部分（或不删除）字符但不改变剩下字符顺序得到的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "coaching", t = "coding"
 * 输出：4
 * 解释：向 s 末尾追加字符串 "ding" ，s = "coachingding" 。
 * 现在，t 是 s ("coachingding") 的一个子序列。
 * 可以证明向 s 末尾追加任何 3 个字符都无法使 t 成为 s 的一个子序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcde", t = "a"
 * 输出：0
 * 解释：t 已经是 s ("abcde") 的一个子序列。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "z", t = "abcde"
 * 输出：5
 * 解释：向 s 末尾追加字符串 "abcde" ，s = "zabcde" 。
 * 现在，t 是 s ("zabcde") 的一个子序列。
 * 可以证明向 s 末尾追加任何 4 个字符都无法使 t 成为 s 的一个子序列。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length, t.length <= 10^5
 * s 和 t 仅由小写英文字母组成
 *
 *
 */
package jzoffer

// @lc code=start
func appendCharacters(s string, t string) int {
	i, n := 0, len(s)
	for j := range t {
		for i < n && s[i] != t[j] {
			i++
		}
		if i == n {
			return len(t) - j
		}
		i++
	}
	return 0
}

// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/append-characters-to-string-to-make-subsequence/solution/tan-xin-pi-pei-by-endlesscheng-d6eq/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func appendCharacters2(s, t string) int {
	j, m := 0, len(t)
	for _, c := range s {
		if byte(c) == t[j] { // s 的字符肯定匹配的是 t 的前缀
			j++
			if j == m {
				return 0
			}
		}
	}
	return m - j
}

// @lc code=end
