/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (44.42%)
 * Likes:    2183
 * Dislikes: 0
 * Total Accepted:    899.6K
 * Total Submissions: 2M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
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

func Len(head *ListNode) int {
	var n int
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

/*
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	num := Len(head)

	pre := &ListNode{Val: 0, Next: head}
	cur := pre
	for i := 0; i < num-n; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	return pre.Next
}

// @lc code=end
