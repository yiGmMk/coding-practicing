/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 *
 * https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (41.58%)
 * Likes:    401
 * Dislikes: 0
 * Total Accepted:    49K
 * Total Submissions: 117.8K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * 给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
 *
 * 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
 *
 * 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * 输出: [1,2],[1,4],[1,6]
 * 解释: 返回序列中的前 3 对数：
 * ⁠    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * 输出: [1,1],[1,1]
 * 解释: 返回序列中的前 2 对数：
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 *
 * 示例 3:
 *
 *
 * 输入: nums1 = [1,2], nums2 = [3], k = 3
 * 输出: [1,3],[2,3]
 * 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums1.length, nums2.length <= 10^5
 * -10^9 <= nums1[i], nums2[i] <= 10^9
 * nums1 和 nums2 均为升序排列
 * 1 <= k <= 1000
 *
 *
 */

// top k的问题首先应该想到用堆解决
package jzoffer

import "container/heap"

// import "container/heap"

// func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
// 	m, n := len(nums1), len(nums2)
// 	h := HP{nil, nums1, nums2}
// 	for i := 0; i < k && i < m; i++ {
// 		h.data = append(h.data, pair{i, 0})
// 	}
// 	for h.Len() > 0 && len(ans) < k {
// 		p := heap.Pop(&h).(pair)
// 		i, j := p.i, p.j
// 		ans = append(ans, []int{nums1[i], nums2[j]})
// 		if j+1 < n {
// 			heap.Push(&h, pair{i, j + 1})
// 		}
// 	}
// 	return
// }

type pair struct{ i, j int }
type HP struct {
	data         []pair
	nums1, nums2 []int
}

func (h HP) Len() int { return len(h.data) }
func (h HP) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h HP) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *HP) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *HP) Pop() interface{} {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

// --------------------------大顶堆---------------------------------------------
/* 维护一个大顶堆,堆顶是最大元素,保存最小的k个元素,
   堆顶是最大(当有k个元素后,出现比堆顶小的元素时将堆顶元素删除,再插入新的元素)
   => 保证堆内是最小的k个元素
*/
var _ heap.Interface = &maxHeap{}

type node struct {
	x, y, sum int
}

// 堆的实现
type maxHeap struct {
	data []node
}

func (h maxHeap) Len() int            { return len(h.data) }
func (h maxHeap) Less(i, j int) bool  { return h.data[i].sum > h.data[j].sum }
func (h maxHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *maxHeap) Push(x interface{}) { h.data = append(h.data, x.(node)) }

// 移除堆顶元素,这里是最大元素
func (h *maxHeap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}

func (h *maxHeap) Top() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.data[0].sum, true
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	out := [][]int{}
	h := maxHeap{}
	for i := 0; i < len(nums1) && i < k; i++ {
		for j := 0; j < len(nums2) && j < k; j++ {
			if h.Len() < k {
				heap.Push(&h, node{x: nums1[i], y: nums2[j], sum: nums1[i] + nums2[j]})
			} else if nums1[i]+nums2[j] <= h.data[0].sum {
				heap.Pop(&h) //删除堆顶元素
				heap.Push(&h, node{x: nums1[i], y: nums2[j], sum: nums1[i] + nums2[j]})
			}
		}
	}
	for h.Len() > 0 && len(out) < k {
		n, _ := heap.Pop(&h).(node)
		out = append(out, []int{n.x, n.y})
	}
	return out
}

// @lc code=end
