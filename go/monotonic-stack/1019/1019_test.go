package jzoffer

import (
	"fmt"
	"testing"
)

func Test1029(t *testing.T) {
	head := build([]int{2, 7, 4, 3, 5})

	res := nextLargerNodes(head)
	fmt.Println(res)
}

func build(data []int) *ListNode {
	p := &ListNode{}
	head := p
	for i, v := range data {
		p.Val = v
		if i == len(data)-1 {
			break
		}
		p.Next = &ListNode{}
		p = p.Next
	}
	return head
}
