package tree

import (
	"reflect"
	"testing"
)

// ABC     A
// BAC    B C
// BCA

func TestGetPostOrderFromPreAndMid(t *testing.T) {
	testcases := []struct {
		preOrder  string
		midOrder  string
		postOrder string
	}{
		{
			preOrder:  "ABC",
			midOrder:  "BAC",
			postOrder: "BCA",
		},
	}

	for _, v := range testcases {
		if got := postTree(v.preOrder, v.midOrder); v.postOrder != got {
			t.Errorf("GetPostOrderFromPreAndMid(%s,%s)=%s, want %s", v.preOrder, v.midOrder, got, v.postOrder)
		}
	}
}

func TestBuildTree(t *testing.T) {
	//  1
	//2   3
	root := buildTree([]int{1, 2, 3}, []int{2, 1, 3})
	if got := postItrate(root); !reflect.DeepEqual(got, []int{2, 3, 1}) {
		t.Errorf("got:%d", got)
	}
}
