package jzoffer

import "testing"

var _3TestCases = []struct {
	s   string
	res int
}{
	//  * 输入: s = "abcabcbb"
	//  * 输出: 3
	//  * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	//  *
	//  *
	//  * 示例 2:
	//  *
	//  *
	//  * 输入: s = "bbbbb"
	//  * 输出: 1
	//  * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
	//  *
	{
		s:   "tmmzuxt",
		res: 5,
	},
	{
		s:   "abba",
		res: 2,
	},
	{
		s:   "abcabcbb",
		res: 3,
	},
	{
		s:   "bbbbb",
		res: 1,
	},
	{
		s:   "pwwkew",
		res: 3,
	},
	{
		s:   " ",
		res: 1,
	},
}

func Test3(t *testing.T) {
	for _, tc := range _3TestCases {
		res := lengthOfLongestSubstring(tc.s)
		if res != tc.res {
			t.Errorf("s:%s,got:%d,expect:%d", tc.s, res, tc.res)
		}

		res = lengthOfLongestSubstringSlidingWindow(tc.s)
		if res != tc.res {
			t.Errorf("s:%s,got:%d,expect:%d", tc.s, res, tc.res)
		}
	}
}
