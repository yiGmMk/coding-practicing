package jzoffer

import (
	"testing"
)

var ts = []struct {
	nums []int
	k    int
	t    int
	res  bool
}{
	{[]int{1, 2, 3, 1}, 3, 0, true},
	{[]int{1, 0, 1, 1}, 1, 2, true},
	{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
}

func Test220(t *testing.T) {
	for _, tc := range ts {
		// 540ms 5.8MB
		res := containsNearbyAlmostDuplicate(tc.nums, tc.k, tc.t)
		if res != tc.res {
			t.Errorf("expected: %v, got: %v, with: %v", tc.res, res, tc)
		}

		// 16ms 6.2MB
		res = containsNearbyAlmostDuplicateBucket(tc.nums, tc.k, tc.t)
		if res != tc.res {
			t.Errorf("expected: %v, got: %v, with: %v", tc.res, res, tc)
		}

		res = containsNearbyAlmostDuplicateSlidingWindow(tc.nums, tc.k, tc.t)
		if res != tc.res {
			t.Errorf("expected: %v, got: %v, with: %v", tc.res, res, tc)
		}
	}
}
