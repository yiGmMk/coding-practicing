/*
 * @lc app=leetcode.cn id=1721 lang=golang
 * @lcpr version=30204
 *
 * [1721] 交换链表中的节点
 */
package golang

// @lcpr-template-start
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapNodes(head *ListNode, k int) *ListNode {
	// 快慢指针 或者 转成数组交换[会超时]
	fast, slow := head, head
	k1, k2 := head, head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
		k1 = fast
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		k2 = slow
	}
	k1.Val, k2.Val = k2.Val, k1.Val
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [7,9,6,6,7,8,3,0,9,5]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n2\n
// @lcpr case=end

*/
