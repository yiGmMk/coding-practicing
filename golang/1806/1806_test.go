package jzoffer

import "testing"

var _1806TestCase = []struct {
	n      int
	expect int
}{
	{
		n:      10,
		expect: 6,
	},
	{
		n:      8,
		expect: 3,
	},
	{
		n:      6,
		expect: 4,
	},
	{
		n:      2,
		expect: 1,
	},
	{
		n:      4,
		expect: 2,
	},
}

func Test1806(t *testing.T) {
	for i, v := range _1806TestCase {
		got := reinitializePermutation(v.n)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
		got = reinitializePermutation1(v.n)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
