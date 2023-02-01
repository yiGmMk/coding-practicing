package jzoffer

import "testing"

var _testcases = []struct {
	key, msg string
	expect   string
}{
	{
		key:    "the quick brown fox jumps over the lazy dog",
		msg:    "vkbs bs t suepuv",
		expect: "this is a secret",
	},
	{
		key:    "eljuxhpwnyrdgtqkviszcfmabo",
		msg:    "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
		expect: "the five boxing wizards jump quickly",
	},
}

func Test2325(t *testing.T) {
	for i, v := range _testcases {
		got := decodeMessage(v.key, v.msg)
		if got != v.expect {
			t.Errorf("[%d].got:%s,expect:%s", i, got, v.expect)
		}
	}
}
