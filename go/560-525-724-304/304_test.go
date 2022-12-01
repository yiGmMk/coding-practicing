package jzoffer

import "testing"

type search struct {
	row1, col1, row2, col2, res int
}

var _304TestCases = []struct {
	nums    [][]int
	searchs []search
}{
	{
		nums: [][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		},
		searchs: []search{
			{
				row1: 2,
				col1: 1,
				row2: 4,
				col2: 3,
				res:  8,
			},
			{
				row1: 1,
				col1: 1,
				row2: 2,
				col2: 2,
				res:  11,
			},
			{
				row1: 1,
				col1: 2,
				row2: 2,
				col2: 4,
				res:  12,
			},
		},
	},
}

func Test304(t *testing.T) {
	for _, tc := range _304TestCases {
		obj := ConstructorOm(tc.nums)
		for _, s := range tc.searchs {
			if res := obj.SumRegionOm(s.row1, s.col1, s.row2, s.col2); res != s.res {
				t.Errorf("expected %d, got %d", s.res, res)
			}
		}
		c := Constructor(tc.nums)
		for _, s := range tc.searchs {
			if res := c.SumRegion(s.row1, s.col1, s.row2, s.col2); res != s.res {
				t.Errorf("expected %d, got %d", s.res, res)
			}
		}
	}
}
