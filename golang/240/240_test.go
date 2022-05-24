package leetcode

import (
	"sort"
	"testing"
)

// 从下往上,从左往右,z字型搜索
var fP = func(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// v=matrix[y][x]
	for y := len(matrix) - 1; y >= 0; y-- {
		for x := 0; x < len(matrix[y]); x++ {
			if y > (len(matrix)-1) || x > (len(matrix[y])-1) {
				return false
			}
			v := matrix[y][x]
			if v == target {
				return true
			} else if v > target {
				y--
				x--
				if x <= 0 {
					x = -1
				}
			}
		}
	}
	return false
}

// 从上往下,从右往左,z型搜
var fN = func(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// v=matrix[y][x]
	for y, row := range matrix {
		for x := len(row) - 1; x >= 0; x-- {
			if y > (len(matrix)-1) || x > (len(matrix[y])-1) {
				return false
			}
			val := matrix[y][x]
			if val == target {
				return true
			} else if val < target {
				y++
				x++
			}
		}
	}
	return false
}

// 每行二分搜索
var fSort = func(matrix [][]int, target int) bool {
	for _, row := range matrix {
		index := sort.SearchInts(row, target)
		if index < len(row) && row[index] == target {
			return true
		}
	}
	return false
}

func TestSearch(t *testing.T) {
	/* 1 2  3
	   4 15 6
	   7 8  9
	*/
	f := fSort // 搜索逻辑代码

	ts := []struct {
		data   [][]int
		target int
		res    bool
	}{
		{
			data:   [][]int{{3}, {6}},
			target: 3,
			res:    true,
		},
		{
			data:   [][]int{{-5, 3}},
			target: 3,
			res:    true,
		},
		{
			data:   [][]int{{-5}},
			target: -5,
			res:    true,
		},
		{
			data: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30}},
			target: 5,
			res:    true,
		},
		{
			data:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			target: 50,
			res:    false,
		},
	}
	for _, v := range ts {
		if f(v.data, v.target) != v.res {
			t.Errorf("error,search failed from %+v,target:%d", v.data, v.target)
		}
	}
}
