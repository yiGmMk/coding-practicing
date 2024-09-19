package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将数组转换为链表
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
