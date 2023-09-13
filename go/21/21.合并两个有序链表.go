/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode.cn/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (66.53%)
 * Likes:    2898
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：l1 = [], l2 = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：l1 = [], l2 = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 两个链表的节点数目范围是 [0, 50]
 * -100
 * l1 和 l2 均按 非递减顺序 排列
 *
 *
 */
package jzoffer

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	p1, p2 := list1, list2
	res := &ListNode{}
	p := res
	for p1 != nil || p2 != nil {
		if p1 != nil && p2 != nil {
			if p1.Val < p2.Val {
				p.Next = &ListNode{Val: p1.Val}
				p1 = p1.Next
			} else {
				p.Next = &ListNode{Val: p2.Val}
				p2 = p2.Next
			}
			p = p.Next
			continue
		}
		if p1 != nil {
			p.Next = &ListNode{Val: p1.Val}
			p = p.Next
			p1 = p1.Next
			continue
		}
		if p2 != nil {
			p.Next = &ListNode{Val: p2.Val}
			p = p.Next
			p2 = p2.Next
			continue
		}
	}
	return res.Next
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
