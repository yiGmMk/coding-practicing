package jzoffer

import "testing"

var _567TestCases = []struct {
	s1, s2 string
	result bool
}{
	{
		s1:     "ab",
		s2:     "eidbaooo",
		result: true,
	},
	{
		s1:     "ab",
		s2:     "eidboaoo",
		result: false,
	},
}

func Test567(t *testing.T) {
	for _, testCase := range _567TestCases {
		if result := checkInclusion(testCase.s1, testCase.s2); result != testCase.result {
			t.Errorf("checkInclusion(%s, %s) return %t, want %t", testCase.s1, testCase.s2, result, testCase.result)
		}
	}
}
