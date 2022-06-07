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
			tree: []string{"1", "null", "0", "null", "null", "0", "1"},
			want: []string{"1", "null", "0", "null", "1"},
		},
	}
	var fix = fDFS
	for _, tc := range ts {
		root, err := util.NewFromArray(tc.tree)
		if err != nil {
			t.Errorf("%v", err)
			continue
		}
		got := fix(root)
		gotArray := util.TreeSource(got)
		if !reflect.DeepEqual(gotArray, tc.want) {
			t.Errorf("\n got: %v \nwant: %v", gotArray, tc.want)
		}
	}
}
