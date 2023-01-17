/*
 * @lc app=leetcode.cn id=1337 lang=golang
 *
 * [1337] 矩阵中战斗力最弱的 K 行
 *
 * https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix/description/
 *
 * algorithms
 * Easy (68.95%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    54.7K
 * Total Submissions: 79.4K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]]\n3'
 *
 * 给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
 *
 * 请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。
 *
 * 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
 *
 * 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：mat =
 * [[1,1,0,0,0],
 * ⁠[1,1,1,1,0],
 * ⁠[1,0,0,0,0],
 * ⁠[1,1,0,0,0],
 * ⁠[1,1,1,1,1]],
 * k = 3
 * 输出：[2,0,3]
 * 解释：
 * 每行中的军人数目：
 * 行 0 -> 2
 * 行 1 -> 4
 * 行 2 -> 1
 * 行 3 -> 2
 * 行 4 -> 5
 * 从最弱到最强对这些行排序后得到 [2,0,3,1,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：mat =
 * [[1,0,0,0],
 * [1,1,1,1],
 * [1,0,0,0],
 * [1,0,0,0]],
 * k = 2
 * 输出：[0,2]
 * 解释：
 * 每行中的军人数目：
 * 行 0 -> 1
 * 行 1 -> 4
 * 行 2 -> 1
 * 行 3 -> 1
 * 从最弱到最强对这些行排序后得到 [0,2,3,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 2
 * 1
 * matrix[i][j] 不是 0 就是 1
 *
 *
 */
package jzoffer

import (
	"container/heap"
	"sort"
)

// @lc code=start

type node struct{ index, cnt int }
type hp []*node

// 小顶堆
func (h *hp) Len() int { return len(*h) }
func (h *hp) Less(i, j int) bool {
	return (*h)[i].cnt < (*h)[j].cnt || ((*h)[i].index < (*h)[j].index && (*h)[i].cnt == (*h)[j].cnt)
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Push(x interface{}) {
	(*h) = append((*h), x.(*node))
}
func (h *hp) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kWeakestRows0(mat [][]int, k int) []int {
	kv := make(map[int]int)
	for i, v := range mat {
		n := sort.Search(len(v), func(i int) bool { return v[i] <= 0 })
		kv[i] = n
		if n < len(v) {
			if v[n] == 1 {
				kv[i]++
			}
		}
	}
	nodes := hp{}
	for k, v := range kv {
		nodes = append(nodes, &node{index: k, cnt: v})
	}
	heap.Init(&nodes)

	res := []int{}
	for k > 0 {
		v := heap.Pop(&nodes)
		if v != nil {
			res = append(res, v.(*node).index)
		}
		k--
	}
	return res
}

// 直接写入堆里
func kWeakestRows(mat [][]int, k int) []int {
	nodes := hp{}
	for i, v := range mat {
		n := sort.Search(len(v), func(i int) bool { return v[i] <= 0 })
		if n < len(v) {
			if v[n] == 1 {
				n++
			}
		}
		nodes = append(nodes, &node{index: i, cnt: n})
	}
	heap.Init(&nodes)

	res := []int{}
	for k > 0 {
		v := heap.Pop(&nodes)
		if v != nil {
			res = append(res, v.(*node).index)
		}
		k--
	}
	return res
}

// @lc code=end
