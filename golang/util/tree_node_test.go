package util

import (
	"reflect"
	"testing"
)

func TestTreeNode(t *testing.T) {
	ts := [][]string{
		[]string{"1", "3", "2", "5", "3", "null", "9"},
		[]string{"1", "2", "3"},
		[]string{}}
	for _, v := range ts {
		tree, err := NewFromArray(v)
		if err != nil {
			t.Errorf("NewFromArray error: %v,source:%s", err, v)
			continue
		}
		if tree == nil && len(v) != 0 {
			t.Errorf("NewFromArray error: %v,source:%s", err, v)
			continue
		}

		source := TreeSource(tree)
		if !reflect.DeepEqual(source, v) {
			t.Errorf("gen tree not right,source_in:%s,source_gen:%s", v, source)
		}
	}
}
