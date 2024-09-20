package main

import (
	"fmt"
	"testing"
)

func Array2LinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	// 创建头节点
	head := &ListNode{Val: arr[0]}
	current := head

	// 遍历数组，逐个创建节点并连接
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

// 打印链表
func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func Test61(t *testing.T) {
	head := Array2LinkedList([]int{2, 3, 4, 5})
	PrintLinkedList(head)

	PrintLinkedList(rotateRight(head, 5))
	head = Array2LinkedList([]int{2, 3, 4, 5})
	PrintLinkedList(rotateRight(head, 2))
}
