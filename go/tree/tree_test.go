package tree

import "testing"

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
