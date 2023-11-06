package jzoffer

import "testing"

func Test2130(t *testing.T) {
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	if got := pairSum(head); got != 6 {
		t.Errorf("got=%d", got)
	}
}
