package jzoffer

import (
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

type TestCase struct {
	tree []string
	want []string
}

func TestPrune(t *testing.T) {
	ts := []TestCase{
		{
			tree: []string{"1", "0", "1", "0", "0", "0", "1"},
			want: []string{"1", "null", "1", "null", "1"},
		},
		{
			tree: []string{"1", "null", "0", "0", "1"},
			want: []string{"1", "null", "0", "null", "1"},
		},
	}
	var fix = fBFS
	for _, tc := range ts {
		root := util.Strs2TreeNode(tc.tree)
		got := fix(root)
		gotArray := util.TreeSource(got)
		if !reflect.DeepEqual(gotArray, tc.want) {
			t.Errorf("\n got: %v \nwant: %v", gotArray, tc.want)
		}
	}
}
