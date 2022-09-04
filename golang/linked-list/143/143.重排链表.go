/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode.cn/problems/reorder-list/description/
 *
 * algorithms
 * Medium (64.29%)
 * Likes:    1018
 * Dislikes: 0
 * Total Accepted:    208.1K
 * Total Submissions: 323.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
 *
 *
 */
package jzoffer

import (
	"github.com/yiGmMk/leetcode/datastructure"
)

type ListNode = datastructure.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
func reorderListSlice(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	ns := []*ListNode{}
	i := 0
	for p := head; p != nil; p = p.Next {
		ns = append(ns, p)
		i++
	}
	n := len(ns)
	i, j := 0, n-1
	for i < j {
		ns[i].Next = ns[j]
		i++
		if i == j {
			break
		}
		ns[j].Next = ns[i]
		j--
	}
	ns[i].Next = nil
	return
}

// 使用切片保存链表
func reorderListSlice2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	ns := []*ListNode{}
	for p := head; p != nil; p = p.Next {
		ns = append(ns, p)
	}
	n, i := len(ns), 0
	for ; i < n/2; i++ {
		ns[i].Next = ns[n-i-1]
		ns[n-i-1].Next = ns[i+1]
	}
	ns[i].Next = nil //最后一个节点的next需要置空,否则通不过=> out of memory: cannot allocate 272629760-byte block
	return
}

// 链表终点+ 链表逆序+ 合并链表

// 1.快慢指针拿到链表的中点
//
//	func middle(head *ListNode) *ListNode { //错误版本
//		s, f := head, head
//		for f != nil {
//			s = s.Next
//			if f.Next.Next == nil {
//				return s
//			}
//			s = f.Next.Next
//		}
//		return s
//	}
func middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 2.链表逆序
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

// 3.合并链表
func merge(l1, l2 *ListNode) *ListNode {
	out := l1
	for l1 != nil && l2 != nil {
		l1next := l1.Next
		l2next := l2.Next

		l1.Next = l2
		l2.Next = l1next

		l1 = l1next
		l2 = l2next
	}
	return out
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	m := middle(head)
	back := m.Next
	m.Next = nil

	rev := reverse(back)

	head = merge(head, rev)
}

// @lc code=end
