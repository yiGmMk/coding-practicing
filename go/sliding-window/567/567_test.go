package jzoffer

import (
	"testing"
)

var _testcase = []struct {
	s1, s2 string
	expect bool
}{
	{
		s1:     "adc",
		s2:     "dcda",
		expect: true,
	},
	{
		s1:     "ab",
		s2:     "eidbaooo",
		expect: true,
	},
	{
		s1:     "ab",
		s2:     "eidboaoo",
		expect: false,
	},
}

func Test5667(t *testing.T) {
	for i, v := range _testcase {
		got := checkInclusion(v.s1, v.s2)
		if got != v.expect {
			t.Errorf("[%d].got:%v,expect:%v", i, got, v.expect)
		}
	}
}
