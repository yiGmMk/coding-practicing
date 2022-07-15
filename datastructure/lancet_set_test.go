package datastructure

import (
	"testing"

	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
)

func TestSet(t *testing.T) {

	set := datastructure.NewSet[int]()
	set.Add(1, 2, 3)

	vs := set.Values()
	if len(vs) != 3 {
		t.Error("should be 3")
	}
}
