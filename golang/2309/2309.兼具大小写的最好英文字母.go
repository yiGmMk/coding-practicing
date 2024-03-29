/*
 * @lc app=leetcode.cn id=2309 lang=golang
 *
 * [2309] 兼具大小写的最好英文字母
 *
 * https://leetcode.cn/problems/greatest-english-letter-in-upper-and-lower-case/description/
 *
 * algorithms
 * Easy (67.23%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    24.8K
 * Total Submissions: 34.5K
 * Testcase Example:  '"lEeTcOdE"'
 *
 * 给你一个由英文字母组成的字符串 s ，请你找出并返回 s 中的 最好
 * 英文字母。返回的字母必须为大写形式。如果不存在满足条件的字母，则返回一个空字符串。
 *
 * 最好 英文字母的大写和小写形式必须 都 在 s 中出现。
 *
 * 英文字母 b 比另一个英文字母 a 更好 的前提是：英文字母表中，b 在 a 之 后 出现。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "lEeTcOdE"
 * 输出："E"
 * 解释：
 * 字母 'E' 是唯一一个大写和小写形式都出现的字母。
 *
 * 示例 2：
 *
 *
 * 输入：s = "arRAzFif"
 * 输出："R"
 * 解释：
 * 字母 'R' 是大写和小写形式都出现的最好英文字母。
 * 注意 'A' 和 'F' 的大写和小写形式也都出现了，但是 'R' 比 'F' 和 'A' 更好。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "AbCdEfGhIjK"
 * 输出：""
 * 解释：
 * 不存在大写和小写形式都出现的字母。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写和大写英文字母组成
 *
 *
 */
package jzoffer

import "unicode"

// @lc code=start
func greatestLetterOn(s string) string {
	res := ""
	rm := make(map[rune]bool)
	update := func(str string) {
		if res == "" {
			res = str
			return
		}
		if str != res && str > res {
			res = str
		}

	}
	for _, v := range s {
		if unicode.IsLower(v) && rm[unicode.ToUpper(v)] {
			update(string(unicode.ToUpper(v)))
		}
		if unicode.IsUpper(v) && rm[unicode.ToLower(v)] {
			update(string(v))
		}
		rm[v] = true
	}
	return res
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/greatest-english-letter-in-upper-and-lower-case/solution/jian-ju-da-xiao-xie-de-zui-hao-ying-wen-o5u2s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func greatestLetter(s string) string {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	for i := 'Z'; i >= 'A'; i-- {
		if set[i] && set[unicode.ToLower(i)] {
			return string(i)
		}
	}
	return ""
}

// @lc code=end
