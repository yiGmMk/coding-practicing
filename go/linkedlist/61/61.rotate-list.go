/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=30204
 *
 * [61] 旋转链表
 */
package main

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := list.New()
	num := 0
	for p := head; p != nil; p = p.Next {
		l.PushBack(p.Val)
		num++
	}
	if k > num {
		k = k % num
	}
	for i := 0; i < k; i++ {
		v := l.Remove(l.Back())
		l.PushFront(v)
	}
	for p, pl := head, l.Front(); p != nil; p = p.Next {
		p.Val = pl.Value.(int)
		pl = pl.Next()
	}
	return head
}

// / 先组成环链,再从尾部移动num-k%num步,断开环链
func rotateRightRing(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n, p := 1, head
	for p.Next != nil {
		p = p.Next
		n++
	}
	p.Next = head // 组成环链
	k %= n
	for i := 1; i <= n-k; i++ {
		p = p.Next
	}
	head, p.Next = p.Next, nil
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/
