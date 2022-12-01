package jzoffer

import (
	"reflect"
	"testing"
)

var _438TestCases = []struct {
	s, p string
	res  []int
}{
	{
		//* 输入: s = "cbaebabacd", p = "abc"
		//* 输出: [0,6]
		s:   "cbaebabacd",
		p:   "abc",
		res: []int{0, 6},
	},
	{
		//* 输入: s = "abab", p = "ab"
		//* 输出: [0,1,2]
		s:   "abab",
		p:   "ab",
		res: []int{0, 1, 2},
	},
}

func Test438(t *testing.T) {
	for _, tc := range _438TestCases {
		res := findAnagrams(tc.s, tc.p)
		if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("got:%+v,expect:%+v", res, tc.res)
		}
	}
}
