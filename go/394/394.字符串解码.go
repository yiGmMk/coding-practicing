/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode.cn/problems/decode-string/description/
 *
 * algorithms
 * Medium (56.91%)
 * Likes:    1625
 * Dislikes: 0
 * Total Accepted:    270.6K
 * Total Submissions: 475.2K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "3[a]2[bc]"
 * 输出："aaabcbc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "3[a2[c]]"
 * 输出："accaccacc"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "2[abc]3[cd]ef"
 * 输出："abcabccdcdcdef"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "abc3[cd]xyz"
 * 输出："abccdcdcdxyz"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 30
 * s 由小写英文字母、数字和方括号 '[]' 组成
 * s 保证是一个 有效 的输入。
 * s 中所有整数的取值范围为 [1, 300]
 *
 *
 */
package jzoffer

import (
	"strconv"
	"strings"
)

// @lc code=start
func decodeStringMy(s string) string {
	var dfs func(s string, i int) (int, string)

	dfs = func(s string, i int) (int, string) {
		var (
			b    strings.Builder
			num  int
			l, r int
		)
		b.Reset()
		for id := i; id < len(s); id++ {
			if s[id]-'0' <= 9 {
				// 解析数字
				nb := []byte{}
				for s[id]-'0' <= 9 {
					nb = append(nb, s[id])
					id++
				}
				num, _ = strconv.Atoi(string(nb))

				// 子串及结束位置
				end, sub := dfs(s, id)
				for j := 0; j < num; j++ {
					b.WriteString(sub)
				}
				id = end
			} else if s[id] == ']' {
				// 括号结束返回子串
				r++
				if l == r || id == len(s)-1 {
					return id, b.String()
				}
			} else if s[id] == '[' {
				l++
			} else {
				b.WriteByte(s[id])
			}
		}
		return l, b.String()
	}
	_, res := dfs(s, 0)
	return res
}

var (
	src string
	ptr int
)

func decodeString(s string) string {
	src = s
	ptr = 0
	return getString()
}

func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigits()
		ptr++
		str := getString()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getString()
}

func getDigits() int {
	ret := 0
	for ; src[ptr]-'0' <= 9; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/decode-string/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
