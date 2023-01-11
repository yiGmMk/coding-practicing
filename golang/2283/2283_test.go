package jzoffer

import "testing"

var _testCases = []struct {
	num    string
	expect bool
}{
	{
		num:    "1210",
		expect: true,
	},
	{
		num:    "030",
		expect: false,
	},
}

func Test2283(t *testing.T) {
	for i, v := range _testCases {
		got := digitCount(v.num)
		if got != v.expect {
			t.Errorf("[%d].got:%v,expect:%v", i, got, v.expect)
		}
		got = digitCount1(v.num)
		if got != v.expect {
			t.Errorf("[%d]1.got:%v,expect:%v", i, got, v.expect)
		}
	}
}
