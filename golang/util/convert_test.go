package util

import (
	"reflect"
	"testing"
)

func TestArrayString(t *testing.T) {
	ts := []struct {
		array string
		res   []string
	}{
		{
			array: "[]",
			res:   []string{},
		},
		{
			array: "[1,2,3]",
			res:   []string{"1", "2", "3"},
		},
		{
			array: "[1,null,2,3]",
			res:   []string{"1", "null", "2", "3"},
		},
	}

	for _, tc := range ts {
		res := LeetcodeArrayStringToArray(tc.array)
		if !reflect.DeepEqual(res, tc.res) {
			t.Errorf("got %v, want %v", res, tc.res)
		}
	}
}
