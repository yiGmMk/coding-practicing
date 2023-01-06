package jzoffer

import (
	"fmt"
	"testing"
)

var _2180TestCase = []struct {
	num    int
	expect int
}{
	{
		num:    4,
		expect: 2,
	},
	{
		num:    30,
		expect: 14,
	},
	{
		num:    1,
		expect: 0,
	},
}

func Test2052(t *testing.T) {

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d & 1 =%d\n", i, i&1)
	}

	for i, v := range _2180TestCase {
		got := countEven1(v.num)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}

		got = countEven(v.num)
		if got != v.expect {
			t.Errorf("[%d]1.got:%d,want:%d", i, got, v.expect)
		}
	}
}
