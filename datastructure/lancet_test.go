package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/slice"
)

func TestLanct(t *testing.T) {
	exist := slice.Contain([]int{1, 2, 3}, 1)
	if !exist {
		t.Error("should be true")
	}
}
