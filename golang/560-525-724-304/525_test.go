package jzoffer

import "testing"

var _525TestCase = []struct {
	nums     []int
	expected int
}{
	{
		nums:     []int{0, 1},
		expected: 2,
	},
	{
		nums:     []int{0, 1, 0},
		expected: 2,
	},
	{
		nums:     []int{0, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		expected: 4,
	},
}

func Test525(t *testing.T) {
	for _, testCase := range _525TestCase {
		actual := findMaxLength(testCase.nums)
		if actual != testCase.expected {
			t.Errorf("\ninput: %v\nactual: %v\nexpected: %v", testCase.nums, actual, testCase.expected)
		}
	}
}
