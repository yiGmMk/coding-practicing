package jzoffer

import "testing"

var _647TestCase = []struct {
	s   string
	res int
}{
	/*
	 * 输入：s = "abc"
	 * 输出：3
	 * 解释：三个回文子串: "a", "b", "c"
	 *
	 *
	 * 示例 2：
	 *
	 *
	 * 输入：s = "aaa"
	 * 输出：6
	 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
	 */
	{
		s:   "abc",
		res: 3,
	},
	{
		s:   "aaa",
		res: 6,
	},
}

func Test647(t *testing.T) {
	for _, tc := range _647TestCase {
		if res := countSubstringsN3(tc.s); res != tc.res {
			t.Errorf("%s,got:%d,want:%d", tc.s, res, tc.res)
		}

		if res := countSubstrings(tc.s); res != tc.res {
			t.Errorf("%s,got:%d,want:%d", tc.s, res, tc.res)
		}
	}
}
