package util

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func TestTrimSapce(t *testing.T) {
	in := []string{"  h e l l   o  ", "  hello  ", "  h e l l o  "}
	expect := "hello"
	for _, s := range in {
		got := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			return r
		}, s)
		if got != expect {
			t.Errorf("trim space error")
		}
	}
}

func TestArrayString(t *testing.T) {
	ts := []struct {
		array string
		res   []string
	}{
		{
			array: "[ 1 , null , 2 , 3 , null , 4 , 5]",
			res:   []string{"1", "null", "2", "3", "null", "4", "5"},
		},
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
