package jzoffer

import (
	"math"
	"testing"
)

var _29TestCases = []struct {
	dividend int
	divisor  int
	result   int
}{
	{
		dividend: math.MaxInt32,
		divisor:  100,
		result:   21474836,
	},
	{
		dividend: 10,
		divisor:  3,
		result:   3,
	},
	{
		dividend: 7,
		divisor:  -3,
		result:   -2,
	},
	{
		dividend: math.MinInt32,
		divisor:  -1,
		result:   math.MaxInt32,
	},
	{
		dividend: math.MinInt32,
		divisor:  1,
		result:   math.MinInt32,
	},
}

func Test29(t *testing.T) {
	for _, testCase := range _29TestCases {
		result := divide(testCase.dividend, testCase.divisor)
		if result != testCase.result {
			t.Errorf("%d / %d = %d, expected %d", testCase.dividend, testCase.divisor, result, testCase.result)
		}

		result = divideBinary(testCase.dividend, testCase.divisor)
		if result != testCase.result {
			t.Errorf("%d / %d = %d, expected %d", testCase.dividend, testCase.divisor, result, testCase.result)
		}
	}
}
