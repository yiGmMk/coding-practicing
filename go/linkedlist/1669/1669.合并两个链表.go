/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 *
 * https://leetcode.cn/problems/merge-in-between-linked-lists/description/
 *
 * algorithms
 * Medium (74.95%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 46.8K
 * Testcase Example:  '[0,1,2,3,4,5]\n3\n4\n[1000000,1000001,1000002]'
 *
 * 给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
 *
 * 请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。
 *
 * 下图中蓝色边和节点展示了操作后的结果：
 *
 * 请你返回结果链表的头指针。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
 * 输出：[0,1,2,1000000,1000001,1000002,5]
 * 解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。
 *
 *
 * 示例 2：
 *
 *
 * 输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 =
 * [1000000,1000001,1000002,1000003,1000004]
 * 输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
 * 解释：上图中蓝色的边和节点为答案链表。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= list1.length <= 10^4
 * 1 <= a <= b < list1.length - 1
 * 1 <= list2.length <= 10^4
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

// 找出a前面一个节点指向list2的头结点
// list2的尾结点后接上b后面一个节点
func mergeInBetween1(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head1 := list1
	var tail2 *ListNode
	for p := list2; p != nil; p = p.Next {
		tail2 = p
	}

	var pre *ListNode // preA
	for i, p := 0, list1; i <= b; i++ {
		if i == a {
			pre.Next = list2
		}
		if i == b {
			tail2.Next = p.Next // b后面的节点
		}
		pre = p
		p = p.Next
	}
	return head1
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	preA, aftB := list1, list1
	for i := 0; i <= b; i++ {
		if i < a-1 {
			preA = preA.Next
		}
		aftB = aftB.Next
	}
	preA.Next = list2
	p := list2
	for p != nil && p.Next != nil {
		p = p.Next
	}
	p.Next = aftB
	return list1
}

// @lc code=end
