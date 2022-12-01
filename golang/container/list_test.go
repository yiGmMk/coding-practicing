package container_test

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	a := list.New()
	a.PushBack(1)
	intersecter := a.PushBack(2)
	a.PushBack(3)

	b := list.New()
	b.PushBack(4)
	b.PushBack(5)
	b.PushBack(intersecter)

	// TODO
	// can't found intersecter
	// find := func(headA, headB *list.Element) *list.Element {
	// 	if headA == nil || headB == nil {
	// 		return nil
	// 	}
	// 	pa, pb := headA, headB
	// 	for pa != pb {
	// 		if pa == nil {
	// 			pa = headB
	// 		} else {
	// 			pa = pa.Next()
	// 		}
	// 		if pb == nil {
	// 			pb = headA
	// 		} else {
	// 			pb = pb.Next()
	// 		}
	// 	}
	// 	return pa
	// }
	// f := find(a.Front(), b.Front())
}

// 1-2-3
// 4-5-2-3
//
// 1-2-3-null- 4  -5-2-3
// 4-5-2- 3  -null-1-2-3
// a+c=m
// b+c=n
// a+c+b=b+c+a
