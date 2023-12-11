/*
 * @lc app=leetcode.cn id=2542 lang=golang
 *
 * [2542] 最大子序列的分数
 *
 * https://leetcode.cn/problems/maximum-subsequence-score/description/
 *
 * algorithms
 * Medium (51.58%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    6.2K
 * Total Submissions: 12K
 * Testcase Example:  '[1,3,3,2]\n[2,1,3,4]\n3'
 *
 * 给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，两者长度都是 n ，再给你一个正整数 k 。你必须从 nums1 中选一个长度为 k
 * 的 子序列 对应的下标。
 *
 * 对于选择的下标 i0 ，i1 ，...， ik - 1 ，你的 分数 定义如下：
 *
 *
 * nums1 中下标对应元素求和，乘以 nums2 中下标对应元素的 最小值 。
 * 用公式表示： (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] ,
 * nums2[i1], ... ,nums2[ik - 1]) 。
 *
 *
 * 请你返回 最大 可能的分数。
 *
 * 一个数组的 子序列 下标是集合 {0, 1, ..., n-1} 中删除若干元素得到的剩余集合，也可以不删除任何元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
 * 输出：12
 * 解释：
 * 四个可能的子序列分数为：
 * - 选择下标 0 ，1 和 2 ，得到分数 (1+3+3) * min(2,1,3) = 7 。
 * - 选择下标 0 ，1 和 3 ，得到分数 (1+3+2) * min(2,1,4) = 6 。
 * - 选择下标 0 ，2 和 3 ，得到分数 (1+3+2) * min(2,3,4) = 12 。
 * - 选择下标 1 ，2 和 3 ，得到分数 (3+3+2) * min(1,3,4) = 8 。
 * 所以最大分数为 12 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
 * 输出：30
 * 解释：
 * 选择下标 2 最优：nums1[2] * nums2[2] = 3 * 10 = 30 是最大可能分数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums1.length == nums2.length
 * 1 <= n <= 10^5
 * 0 <= nums1[i], nums2[j] <= 10^5
 * 1 <= k <= n
 *
 *
 */
package jzoffer

import (
	"container/heap"
	"sort"
)

// @lc code=start
// TODO: https://leetcode.cn/problems/maximum-subsequence-score/?envType=study-plan-v2&envId=leetcode-75
func maxScore(nums1, nums2 []int, k int) int64 {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].y > a[j].y })

	h := hp{} // 复用内存
	for _, p := range a[:k] {
		sum += p.x
		h.IntSlice = append(h.IntSlice, p.x)
	}
	ans := sum * a[k-1].y
	heap.Init(&h)
	for _, p := range a[k:] {
		if p.x > h.IntSlice[0] {
			x, ok := heap.Pop(&h).(int)
			if !ok {
				continue
			}
			sum += p.x - x
			heap.Push(&h, p.x)
			ans = max(ans, sum*p.y)
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (h *hp) Pop() interface{} {
	top := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return top
}
func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// func (h hp) replace(v int) int {
// 	top := h.IntSlice[0]
// 	h.IntSlice[0] = v
// 	heap.Fix(&h, 0)
// 	return top
// }

// @lc code=end
