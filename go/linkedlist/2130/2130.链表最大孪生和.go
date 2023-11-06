/*
 * @lc app=leetcode.cn id=2130 lang=golang
 *
 * [2130] 链表最大孪生和
 *
 * https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/description/
 *
 * algorithms
 * Medium (78.86%)
 * Likes:    48
 * Dislikes: 0
 * Total Accepted:    18.2K
 * Total Submissions: 23.2K
 * Testcase Example:  '[5,4,2,1]'
 *
 * 在一个大小为 n 且 n 为 偶数 的链表中，对于 0 <= i <= (n / 2) - 1 的 i ，第 i 个节点（下标从 0
 * 开始）的孪生节点为第 (n-1-i) 个节点 。
 *
 *
 * 比方说，n = 4 那么节点 0 是节点 3 的孪生节点，节点 1 是节点 2 的孪生节点。这是长度为 n = 4 的链表中所有的孪生节点。
 *
 *
 * 孪生和 定义为一个节点和它孪生节点两者值之和。
 *
 * 给你一个长度为偶数的链表的头节点 head ，请你返回链表的 最大孪生和 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：head = [5,4,2,1]
 * 输出：6
 * 解释：
 * 节点 0 和节点 1 分别是节点 3 和 2 的孪生节点。孪生和都为 6 。
 * 链表中没有其他孪生节点。
 * 所以，链表的最大孪生和是 6 。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：head = [4,2,2,3]
 * 输出：7
 * 解释：
 * 链表中的孪生节点为：
 * - 节点 0 是节点 3 的孪生节点，孪生和为 4 + 3 = 7 。
 * - 节点 1 是节点 2 的孪生节点，孪生和为 2 + 2 = 4 。
 * 所以，最大孪生和为 max(7, 4) = 7 。
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：head = [1,100000]
 * 输出：100001
 * 解释：
 * 链表中只有一对孪生节点，孪生和为 1 + 100000 = 100001 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表的节点数目是 [2, 10^5] 中的 偶数 。
 * 1 <= Node.val <= 10^5
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
// O(n) 存储空间
func pairSumMy(head *ListNode) (res int) {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next

	}
	n := len(nums)
	for i := 0; i <= len(nums)/2; i++ {
		res = max(res, nums[i]+nums[n-1-i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pairSum(head *ListNode) (res int) {
	var (
		fast, slow = head, head
	)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	back := reverse(slow)
	for back != nil {
		res = max(res, head.Val+back.Val)
		head = head.Next
		back = back.Next
	}
	return
}

func reverse(head *ListNode) *ListNode {
	var (
		pre *ListNode
		cur = head
	)
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	return pre
}

// @lc code=end
