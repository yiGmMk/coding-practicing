/*
 * @lc app=leetcode.cn id=2487 lang=golang
 *
 * [2487] 从链表中移除节点
 *
 * https://leetcode.cn/problems/remove-nodes-from-linked-list/description/
 *
 * algorithms
 * Medium (69.38%)
 * Likes:    14
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 9K
 * Testcase Example:  '[5,2,13,3,8]'
 *
 * 给你一个链表的头节点 head 。
 *
 * 对于列表中的每个节点 node ，如果其右侧存在一个具有 严格更大 值的节点，则移除 node 。
 *
 * 返回修改后链表的头节点 head 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [5,2,13,3,8]
 * 输出：[13,8]
 * 解释：需要移除的节点是 5 ，2 和 3 。
 * - 节点 13 在节点 5 右侧。
 * - 节点 13 在节点 2 右侧。
 * - 节点 8 在节点 3 右侧。
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,1,1,1]
 * 输出：[1,1,1,1]
 * 解释：每个节点的值都是 1 ，所以没有需要移除的节点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定列表中的节点数目在范围 [1, 10^5] 内
 * 1 <= Node.val <= 10^5
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
func removeNodes(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	for head != nil {
		if len(nodes) > 0 {
			j, big := len(nodes)-1, false
			if head.Val > nodes[j].Val {
				for i := len(nodes) - 1; i >= 0; i-- {
					if head.Val > nodes[i].Val {
						j = i
						big = true
					}
				}
			}
			if big {
				if j == 0 {
					nodes = make([]*ListNode, 0, 0)
				} else {
					// new := make([]*ListNode, j, j)
					// copy(new, )
					nodes = nodes[:j]
				}
			}
		}
		nodes = append(nodes, head)
		head = head.Next
	}

	// fmt.Println(len(nodes), cap(nodes), nodes)

	for i := 0; i < len(nodes); i++ {
		if i == len(nodes)-1 {
			nodes[i].Next = nil
		} else {
			nodes[i].Next = nodes[i+1]
		}
	}
	return nodes[0]
}

// https://leetcode.cn/problems/remove-nodes-from-linked-list/solution/di-gui-jian-ji-xie-fa-by-endlesscheng-jfwi/
func removeNodesRecursion(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	node := removeNodesRecursion(head.Next)
	if node.Val > head.Val {
		return node //删除head节点
	}
	head.Next = node // 不删除
	return head
}

// @lc code=end
