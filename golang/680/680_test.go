package jzoffer

import "testing"

var _680TestCases = []struct {
	s  string
	is bool
}{
	// * 输入: s = "aba"
	// * 输出: true
	// *
	// *
	// * 示例 2:
	// *
	// *
	// * 输入: s = "abca"
	// * 输出: true
	// * 解释: 你可以删除c字符。
	// *
	// *
	// * 示例 3:
	// *
	// *
	// * 输入: s = "abc"
	// * 输出: false
	{
		s:  "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
		is: true,
	},
	{
		s:  "aba",
		is: true,
	},
	{
		s:  "abca",
		is: true,
	},
	{
		s:  "abc",
		is: false,
	},
}

/* abcdcdb*/
func Test680(t *testing.T) {
	for _, tc := range _680TestCases {
		if res := validPalindrome1(tc.s); res != tc.is {
			t.Errorf("%s,got:%+v,expect:%+v", tc.s, res, tc.is)
		}

		if res := validPalindrome(tc.s); res != tc.is {
			t.Errorf("%s,got:%+v,expect:%+v", tc.s, res, tc.is)
		}
	}
}
