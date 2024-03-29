/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (52.39%)
 * Likes:    1495
 * Dislikes: 0
 * Total Accepted:    496.3K
 * Total Submissions: 946.8K
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */
package jzoffer

import "github.com/yiGmMk/leetcode/datastructure"

type ListNode = datastructure.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 找出链表中点,后半部分翻转下,后面就和双指针判断字符串是否是回文差不多了
func isPalindrome(head *ListNode) bool {
	middle := func(node *ListNode) *ListNode {
		var slow, fast *ListNode = node, node
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	reverse := func(node *ListNode) *ListNode {
		var pre, cur *ListNode = nil, node
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}

	mid := middle(head)
	rev := reverse(mid)
	ph, pr := head, rev
	for ph != nil && pr != nil {
		if ph.Val != pr.Val {
			return false
		}
		ph = ph.Next
		pr = pr.Next
	}
	return true
}

// @lc code=end
