/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 *
 * https://leetcode.cn/problems/next-greater-node-in-linked-list/description/
 *
 * algorithms
 * Medium (60.98%)
 * Likes:    301
 * Dislikes: 0
 * Total Accepted:    51.7K
 * Total Submissions: 80.4K
 * Testcase Example:  '[2,1,5]'
 *
 * 给定一个长度为 n 的链表 head
 *
 * 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
 *
 * 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i
 * 个节点没有下一个更大的节点，设置 answer[i] = 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [2,1,5]
 * 输出：[5,5,0]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [2,7,4,3,5]
 * 输出：[7,0,5,5,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数为 n
 * 1 <= n <= 10^4
 * 1 <= Node.val <= 10^9
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
// 单调栈,没有找到下一个更大值的元素
func nextLargerNodes(head *ListNode) []int {
	var (
		res   []int
		stack []struct{ v, i int }
		cur   = head
		index = -1
	)

	for cur != nil {
		index++
		res = append(res, 0)

		for len(stack) > 0 && stack[len(stack)-1].v < cur.Val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top.i] = cur.Val
		}

		stack = append(stack, struct {
			v int
			i int
		}{v: cur.Val, i: index})
		cur = cur.Next
	}
	return res
}

// @lc code=end
