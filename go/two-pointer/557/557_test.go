package jzoffer

import "testing"

var _testcases = []struct {
	str, expect string
}{
	{
		str:    "Let's take LeetCode contest",
		expect: "s'teL ekat edoCteeL tsetnoc",
	},
	{
		str:    "God Ding",
		expect: "doG gniD",
	},
}

func Test557(t *testing.T) {
	for i, v := range _testcases {
		got := reverseWords(v.str)
		if got != v.expect {
			t.Errorf("[%d]1.got:%v,expect:%v", i, got, v.expect)
		}
	}
}
