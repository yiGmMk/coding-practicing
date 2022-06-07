package jzoffer

import (
	"reflect"
	"testing"

	"github.com/yiGmMk/leetcode/golang/util"
)

var fBFS = rightSideView

type TestCase struct {
	tree []string
	res  []int
}

func TestSolution(t *testing.T) {
	ts := []TestCase{
		{
			tree: []string{"1", "2", "3", "null", "5", "null", "4"},
			res:  []int{1, 3, 4},
		},
		{
			tree: []string{"1", "null", "3"},
			res:  []int{1, 3},
		},
	}
	for _, tc := range ts {
		root, err := util.NewFromArray(tc.tree)
		if err != nil {
			t.Errorf("生成tree失败,输入参数:%s", tc.tree)
		}
		res := fBFS(root)
		if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("error,case:%+v,res:%+v", tc, res)
		}
	}
}
