package jzoffer

import "testing"

var _2052TestCase = []struct {
	num    []int
	expect int
}{
	{
		num:    []int{30},
		expect: 3,
	},
	{
		num:    []int{2, 4, 3, 7, 10, 6},
		expect: 4,
	},
	{
		num:    []int{2, 4, 8, 16},
		expect: 1,
	},
}

func Test2052(t *testing.T) {
	for i, v := range _2052TestCase {
		got := distinctPrimeFactors(v.num)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = distinctPrimeFactors1(v.num)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,want:%d", i, got, v.expect)
		}
	}
}
