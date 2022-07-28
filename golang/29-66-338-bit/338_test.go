package jzoffer

import (
	"reflect"
	"testing"
)

var _338TestCases = []struct {
	n      int
	result []int
}{
	{
		n:      2,
		result: []int{0, 1, 1},
	},
	{
		n:      5,
		result: []int{0, 1, 1, 2, 1, 2}, // 0 1 10 11 100 101
	},
}

func Test338(t *testing.T) {
	check := func(n int, result, expect []int, t *testing.T) {
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("%d 's bitcount: %v, expected %v", n, result, expect)
		}
	}
	for _, testCase := range _338TestCases {
		result := countBits(testCase.n)
		check(testCase.n, result, testCase.result, t)

		result = countBitsOn(testCase.n)
		check(testCase.n, result, testCase.result, t)

		result = countBitsOn(testCase.n)
		check(testCase.n, result, testCase.result, t)

		result = countBitsLowBit(testCase.n)
		check(testCase.n, result, testCase.result, t)

		result = countBitsBrianKernighan(testCase.n)
		check(testCase.n, result, testCase.result, t)
	}
}
