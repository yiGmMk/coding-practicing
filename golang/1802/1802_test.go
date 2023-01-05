package jzoffer

import "testing"

var _1802TestCase = []struct {
	n, index, maxSum int
	expect           int
}{
	{
		n:      4,
		index:  2,
		maxSum: 6,
		expect: 2,
	},
	{
		n:      6,
		index:  1,
		maxSum: 10,
		expect: 3,
	},
	{
		n:      1,
		index:  0,
		maxSum: 24,
		expect: 24,
	},
}

func Test1802(t *testing.T) {
	for i, v := range _1802TestCase {
		got := maxValue(v.n, v.index, v.maxSum)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
		got = maxValue0(v.n, v.index, v.maxSum)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
		got = maxValue00(v.n, v.index, v.maxSum)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = maxValue1(v.n, v.index, v.maxSum)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = maxValue2(v.n, v.index, v.maxSum)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
