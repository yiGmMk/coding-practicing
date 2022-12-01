package jzoffer

import (
	"testing"
)

var _125TestCases = []struct {
	s  string
	is bool
}{
	// * 输入: "A man, a plan, a canal: Panama"
	// * 输出: true
	// * 解释："amanaplanacanalpanama" 是回文串
	// *
	// *
	// * 示例 2:
	// *
	// *
	// * 输入: "race a car"
	// * 输出: false
	// * 解释："raceacar" 不是回文串
	{
		s:  "",
		is: true,
	},
	{
		s:  "A man, a plan, a canal: Panama",
		is: true,
	},
	{
		s:  "race a car",
		is: false,
	},
	{
		s:  ".,",
		is: true,
	},
}

func Test125(t *testing.T) {
	println(isPalindrome1(_125TestCases[1].s))
	for _, tc := range _125TestCases {
		if res := isPalindrome(tc.s); res != tc.is {
			t.Errorf("%s,got:%+v,expect:%+v", tc.s, res, tc.is)
		}

		if res := isPalindrome1(tc.s); res != tc.is {
			t.Errorf("%s,got:%+v,expect:%+v", tc.s, res, tc.is)
		}

		if res := isPalindrome2(tc.s); res != tc.is {
			t.Errorf("%s,got:%+v,expect:%+v", tc.s, res, tc.is)
		}

		if res := isPalindrome3(tc.s); res != tc.is {
			t.Errorf("%s,got:%+v,expect:%+v", tc.s, res, tc.is)
		}
	}
}
