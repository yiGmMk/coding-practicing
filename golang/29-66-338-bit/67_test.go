package jzoffer

import (
	"fmt"
	"testing"
)

var _67TestCases = []struct {
	a, b, res string
}{
	{
		a:   "1111",
		b:   "1111",
		res: "11110",
	},
	{
		a:   "11",
		b:   "1",
		res: "100",
	},
	{
		a:   "1010",
		b:   "1011",
		res: "10101",
	},
}

func Test167(t *testing.T) {
	res := make([]int, 0, 100)
	fmt.Println(res, len(res), cap(res))
	for _, tc := range _67TestCases {
		if res := addBinary(tc.a, tc.b); res != tc.res {
			t.Errorf("addBinary(%s, %s) => %s, want %s", tc.a, tc.b, res, tc.res)
		}

		if res := addBinary2(tc.a, tc.b); res != tc.res {
			t.Errorf("addBinary2(%s, %s) => %s, want %s", tc.a, tc.b, res, tc.res)
		}
	}
}
