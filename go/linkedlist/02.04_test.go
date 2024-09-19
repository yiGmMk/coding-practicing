package linkedlist

import (
	"testing"
)

func Test0204(t *testing.T) {
	head := Array2LinkedList([]int{1, 4, 3, 2, 5, 2})
	PrintLinkedList(head)
	PrintLinkedList(partition(head, 3))
}

func Test02042(t *testing.T) {
	head := Array2LinkedList([]int{1, 4, 3, 2, 5, 2})
	PrintLinkedList(head)
	PrintLinkedList(partition2(head, 3))
}
