package jzoffer

import "testing"

var _testcase = []struct {
	s      string
	expect int
}{
	{
		s:      "abcabcbb",
		expect: 3,
	},
	{
		s:      "tmmzuxt",
		expect: 5,
	},
	{
		s:      "pwwkew",
		expect: 3,
	},
	{
		s:      "aab",
		expect: 2,
	},
	{
		s:      "abba",
		expect: 2,
	},

	{
		s:      "dvdf",
		expect: 3,
	},
	{},
}

func Test3(t *testing.T) {
	for i, v := range _testcase {
		got := lengthOfLongestSubstring(v.s)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
