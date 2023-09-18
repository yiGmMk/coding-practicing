package jzoffer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test238(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3, 4})
	if d := cmp.Diff(res, []int{24, 12, 8, 6}); d != "" {
		t.Errorf("%v,%s", res, d)
	}
}
