package jzoffer

import "testing"

var _2486TestCase = []struct {
	s, t   string
	expect int
}{
	{
		s:      "coaching",
		t:      "coding",
		expect: 4,
	},
	{
		s:      "abcde",
		t:      "a",
		expect: 0,
	},
	{
		s:      "z",
		t:      "abcde",
		expect: 5,
	},
}

func Test2486(t *testing.T) {
	for i, v := range _2486TestCase {
		got := appendCharacters(v.s, v.t)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d\n", i, got, v.expect)
		}
	}
}
