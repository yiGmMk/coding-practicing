package jzoffer

import "testing"

var _testCase = []struct {
	s         string
	knowledge [][]string
	expect    string
}{
	{
		s:         "(name)is(age)yearsold",
		knowledge: [][]string{{"name", "bob"}, {"age", "two"}},
		expect:    "bobistwoyearsold",
	},
}

func Test1807(t *testing.T) {
	for i, v := range _testCase {
		got := evaluate(v.s, v.knowledge)
		if got != v.expect {
			t.Errorf("[%d].got:%s,expect:%s", i, got, v.expect)
		}
	}
}
