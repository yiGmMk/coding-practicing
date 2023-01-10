package jzoffer

import "testing"

var _testCases = []struct {
	n, expect int
}{
	{
		n:      1,
		expect: 1,
	},
	{
		n:      5,
		expect: 2,
	},

	{
		n:      3,
		expect: 2,
	},
	{
		n:      8,
		expect: 3,
	},
}

func Test441(t *testing.T) {
	for i, v := range _testCases {
		got := arrangeCoins(v.n)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
