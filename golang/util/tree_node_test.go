package util

import (
	"reflect"
	"testing"
)

func TestTreeNode(t *testing.T) {
	ts := [][]string{
		[]string{"1", "2", "3", "null", "null", "4", "5", "6", "7"},
		[]string{"1", "2", "3", "null", "null", "4", "5"},
		[]string{"1", "3", "2", "5", "3", "null", "9"},
		[]string{"1", "2", "3"},
		[]string{}}
	for _, v := range ts {
		tree := Strs2TreeNode(v)
		if tree == nil && len(v) != 0 {
			continue
		}

		source := TreeSource(tree)
		if !reflect.DeepEqual(source, v) {
			t.Errorf("gen tree not right,source_in:%s,source_gen:%s", v, source)
		}

		source = Tree2Strs(tree)
		if !reflect.DeepEqual(source, v) {
			t.Errorf("gen tree not right,source_in:%s,source_gen:%s", v, source)
		}
	}
}
