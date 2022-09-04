/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode.cn/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (60.07%)
 * Likes:    550
 * Dislikes: 0
 * Total Accepted:    109.8K
 * Total Submissions: 182.8K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入：l1 = [7,2,4,3], l2 = [5,6,4]
 * 输出：[7,8,0,7]
 *
 *
 * 示例2：
 *
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[8,0,7]
 *
 *
 * 示例3：
 *
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 100]
 * 0 <= node.val <= 9
 * 输入数据保证链表代表的数字无前导 0
 *
 *
 *
 *
 * 进阶：如果输入链表不能翻转该如何解决？
 *
 */
package jzoffer

import (
	"container/list"

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

// -------------------1. 翻转链表实现-----------------------------------
func addTwoNumbersReverse(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := reverse(l1), reverse(l2)
	carry := 0
	res := &ListNode{}
	out := res
	add := func(n1, n2 *ListNode) int {
		if n1 == nil && n2 == nil {
			return 0
		}
		var v int
		if n1 != nil {
			v += n1.Val
		}
		if n2 != nil {
			v += n2.Val
		}
		if carry > 0 {
			v += 1
			carry -= 1
		}
		if v >= 10 {
			carry += v / 10
		}
		res.Val = v % 10
		if (n1 != nil && n1.Next != nil) || (n2 != nil && n2.Next != nil) {
			res.Next = &ListNode{}
			res = res.Next
		}
		return v
	}
	for p1 != nil && p2 != nil {
		add(p1, p2)
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		add(p1, nil)
		p1 = p1.Next
	}
	for p2 != nil {
		add(nil, p2)
		p2 = p2.Next
	}

	for carry > 0 {
		v := carry % 10
		carry /= 10
		res.Next = &ListNode{Val: v}
		res = res.Next
	}
	node := reverse(out) // 都已经翻转了,这里的指向可以优化,少一次翻转
	return node
}

func reverse(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	var pre, cur *ListNode = nil, l
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 1.1 优化通过翻转的实现,优化结果链表的构建,少一次翻转
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := reverse(l1), reverse(l2)
	carry := 0
	var res *ListNode

	for p1 != nil || p2 != nil || carry > 0 {
		var v int
		if p1 != nil {
			v += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v += p2.Val
			p2 = p2.Next
		}
		v += carry
		carry = v / 10
		v = v % 10
		node := &ListNode{Val: v, Next: res}
		res = node
	}
	return res
}

// ----------------------- 2.不翻转,使用stack-----------------------------
func addTwoNumbersStack(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := list.New(), list.New()
	for l1 != nil {
		s1.PushBack(l1)
		l1 = l1.Next
	}
	for l2 != nil {
		s2.PushBack(l2)
		l2 = l2.Next
	}
	var carry int
	var res *ListNode
	for s1.Len() > 0 || s2.Len() > 0 || carry > 0 {
		var v int
		if s1.Len() > 0 {
			node := s1.Back()
			v += node.Value.(*ListNode).Val
			s1.Remove(node)
		}
		if s2.Len() > 0 {
			node := s2.Back()
			v += node.Value.(*ListNode).Val
			s2.Remove(node)
		}
		v += carry
		carry = v / 10
		v = v % 10

		node := &ListNode{Val: v, Next: res}
		res = node
	}
	return res
}

// @lc code=end
