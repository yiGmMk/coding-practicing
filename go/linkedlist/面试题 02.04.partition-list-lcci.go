/*
 * @lc app=leetcode.cn id=面试题 02.04 lang=golang
 * @lcpr version=30204
 *
 * [面试题 02.04] 分割链表
 */
package linkedlist

import (
	"container/list"
)

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
// 多一倍内存 O(n)
func partition(head *ListNode, x int) *ListNode {
	l := list.New()
	// 先找>=x的节点xlarge,当找到后,小于x的插到xLarge前面
	var xLarge *list.Element
	for p := head; p != nil; p = p.Next {
		if p.Val < x && xLarge != nil {
			xLarge = l.InsertBefore(p, xLarge)
			continue
		}
		v := l.PushBack(p)
		if p.Val >= x && xLarge == nil {
			xLarge = v
		}
	}
	res := &ListNode{}
	head = res
	for p := l.Front(); p != nil; p = p.Next() {
		res.Next = &ListNode{Val: p.Value.(*ListNode).Val}
		res = res.Next
	}
	return head.Next
}

// bug
// 原始值 1 -> 4 -> 3 -> 2 -> 5 -> 2 -> nil ,x =3
// 结果 1 -> 2 -> 2 -> 4 -> 3 -> 5 -> nil

func partition2(head *ListNode, x int) *ListNode {
	l := list.New()
	// 先找>=x的节点xlarge,当找到后,小于x的插到xLarge前面
	var xLarge *list.Element
	for p := head; p != nil; p = p.Next {
		if p.Val < x && xLarge != nil {
			xLarge = l.InsertBefore(p, xLarge)
			continue
		}
		v := l.PushBack(p)
		if p.Val >= x && xLarge == nil {
			xLarge = v
		}
	}
	for p, p1 := l.Front(), head; p != nil; p = p.Next() {
		p1.Val = p.Value.(*ListNode).Val
		p1 = p1.Next
	}
	return head
}

/*bug fix,往list插入的原链表的指针,第二个for循环修改了指针的值,导致部分值拷贝错误 */
func partition2fix(head *ListNode, x int) *ListNode {
	l := list.New()
	// 先找>=x的节点xlarge,当找到后,小于x的插到xLarge前面
	var xLarge *list.Element
	for p := head; p != nil; p = p.Next {
		if p.Val < x && xLarge != nil {
			xLarge = l.InsertBefore(p.Val, xLarge)
			continue
		}
		v := l.PushBack(p.Val)
		if p.Val >= x && xLarge == nil {
			xLarge = v
		}
	}
	for p, p1 := l.Front(), head; p != nil; p = p.Next() {
		p1.Val = p.Value.(int)
		p1 = p1.Next
	}
	return head
}

// 通过2个指针实现
func partition3(head *ListNode, x int) *ListNode {
	// 迭代实现
	// 分隔链表，所有小于x的节点都出现在大于或等于x的节点前，相对位置不变
	// 虚拟头结点，创建2个虚拟头结点，分left right，最后合并
	// 时间O(n)，空间O(1)
	lDummy, rDummy := &ListNode{}, &ListNode{}
	lIdx, rIdx := lDummy, rDummy
	cur := head
	// 遍历链表，小于x的放入lIdx链表，大于等于放rIdx
	for cur != nil {
		if cur.Val < x {
			lIdx.Next = cur
			lIdx = lIdx.Next
		} else {
			rIdx.Next = cur
			rIdx = rIdx.Next
		}
		cur = cur.Next
	}
	// rIdx最后指向空
	// lIdx的最后节点指向rIdx的第一个节点
	rIdx.Next = nil
	lIdx.Next = rDummy.Next
	return lDummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,4,3,2,5,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n2\n
// @lcpr case=end

*/
