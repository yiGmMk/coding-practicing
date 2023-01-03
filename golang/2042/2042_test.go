package jzoffer

import "testing"

var _2042TestCase = []struct {
	s      string
	expect bool
}{
	{
		s:      "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
		expect: true,
	},
	{
		s:      "hello world 5 x 5",
		expect: false,
	},
	{
		s:      "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
		expect: false,
	},
	{
		s:      "4 5 11 26",
		expect: true,
	},
}

func Test2042(t *testing.T) {
	for i, v := range _2042TestCase {
		got := areNumbersAscending(v.s)
		if got != v.expect {
			t.Errorf("[%d].got:%v,want:%v", i, got, v.expect)
		}
	}
}
