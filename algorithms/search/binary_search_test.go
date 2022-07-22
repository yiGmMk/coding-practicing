package algorithms

import (
	"log"
	"sort"
	"testing"
)

var _searchTestCases = []SearchTestCase{
	{
		values:   []int{},
		target:   100,
		expected: -1,
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:   1,
		expected: 0,
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:   100,
		expected: -1,
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:   5,
		expected: 4,
	},
}

var _searchStringTestCases = []SearchStringTestCase{
	{
		values:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		target:   "e",
		expected: 4,
	},
	{
		values:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		target:   "z",
		expected: -1,
	},
}

func TestBinary(t *testing.T) {
	for _, val := range _searchTestCases {
		index := BinarySearch(val.values, val.target)
		if index != val.expected {
			t.Errorf("BinarySearch(%v, %v) = %v, expected %v",
				val.values, val.target, index, val.expected)
		}

		insertIndex := sort.SearchInts(val.values, val.target)
		if index != insertIndex {
			log.Println(
				"search ", val.target,
				"in ", val.values,
				"my_version:", index, "sort.SearchInts:", insertIndex)
		}
	}
}

func TestBinaryString(t *testing.T) {
	for _, val := range _searchStringTestCases {
		index := BinarySearch(val.values, val.target)
		if index != val.expected {
			t.Errorf("BinarySearch(%v, %v) = %v, expected %v",
				val.values, val.target, index, val.expected)
		}

		insertIndex := sort.SearchStrings(val.values, val.target)
		if index != insertIndex {
			log.Println(
				"search ", val.target,
				"in ", val.values,
				"my_version:", index, "sort.SearchInts:", insertIndex)
		}
	}
}
