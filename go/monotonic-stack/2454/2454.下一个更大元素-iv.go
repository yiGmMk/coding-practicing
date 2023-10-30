/*
 * @lc app=leetcode.cn id=2454 lang=golang
 *
 * [2454] 下一个更大元素 IV
 *
 * https://leetcode.cn/problems/next-greater-element-iv/description/
 *
 * algorithms
 * Hard (48.94%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    4.6K
 * Total Submissions: 9.3K
 * Testcase Example:  '[2,4,0,9,6]'
 *
 * 给你一个下标从 0 开始的非负整数数组 nums 。对于 nums 中每一个整数，你必须找到对应元素的 第二大 整数。
 *
 * 如果 nums[j] 满足以下条件，那么我们称它为 nums[i] 的 第二大 整数：
 *
 *
 * j > i
 * nums[j] > nums[i]
 * 恰好存在 一个 k 满足 i < k < j 且 nums[k] > nums[i] 。
 *
 *
 * 如果不存在 nums[j] ，那么第二大整数为 -1 。
 *
 *
 * 比方说，数组 [1, 2, 4, 3] 中，1 的第二大整数是 4 ，2 的第二大整数是 3 ，3 和 4 的第二大整数是 -1 。
 *
 *
 * 请你返回一个整数数组 answer ，其中 answer[i]是 nums[i] 的第二大整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,4,0,9,6]
 * 输出：[9,6,6,-1,-1]
 * 解释：
 * 下标为 0 处：2 的右边，4 是大于 2 的第一个整数，9 是第二个大于 2 的整数。
 * 下标为 1 处：4 的右边，9 是大于 4 的第一个整数，6 是第二个大于 4 的整数。
 * 下标为 2 处：0 的右边，9 是大于 0 的第一个整数，6 是第二个大于 0 的整数。
 * 下标为 3 处：右边不存在大于 9 的整数，所以第二大整数为 -1 。
 * 下标为 4 处：右边不存在大于 6 的整数，所以第二大整数为 -1 。
 * 所以我们返回 [9,6,6,-1,-1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,3]
 * 输出：[-1,-1]
 * 解释：
 * 由于每个数右边都没有更大的数，所以我们返回 [-1,-1] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^9
 *
 *
 */

package jzoffer

import "container/heap"

// @lc code=start
type (
	minHeap []node
	node    struct{ i, v int }
)

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i].v <= (*h)[j].v
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *minHeap) Push(v any) {
	*h = append(*h, v.(node))
}

// 单调栈+小顶堆
func secondGreaterElement(nums []int) (res []int) {
	var (
		stack []node
		h     minHeap
	)

	res = make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	for i, v := range nums {
		// 在小顶堆中的已有一个大于的数,此时再找的即是第二个大于的数
		for h.Len() > 0 && v > h[0].v {
			top, _ := heap.Pop(&h).(node)
			res[top.i] = v
		}
		// 单调栈,单调递减,找出第一个大于的数推入堆中
		for len(stack) > 0 && v > stack[len(stack)-1].v {
			top := stack[len(stack)-1]
			heap.Push(&h, top)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, node{i: i, v: v})
	}

	return
}

// @lc code=end
