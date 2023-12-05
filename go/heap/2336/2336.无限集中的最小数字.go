/*
* @lc app=leetcode.cn id=2336 lang=golang
*
* [2336] 无限集中的最小数字
*
* https://leetcode.cn/problems/smallest-number-in-infinite-set/description/
*
  - algorithms
  - Medium (72.79%)
  - Likes:    70
  - Dislikes: 0
  - Total Accepted:    31.8K
  - Total Submissions: 43.7K
  - Testcase Example:  '["SmallestInfiniteSet","addBack","popSmallest","popSmallest","popSmallest","addBack","popSmallest","popSmallest","popSmallest"]\n' +
    '[[],[2],[],[],[],[1],[],[],[]]'

*
* 现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。
*
* 实现 SmallestInfiniteSet 类：
*
*
* SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
* int popSmallest() 移除 并返回该无限集中的最小整数。
* void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集最后。
*
*
*
*
* 示例：
*
*
* 输入
* ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest",
* "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
* [[], [2], [], [], [], [1], [], [], []]
* 输出
* [null, null, 1, 2, 3, null, 1, 4, 5]
*
* 解释
* SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
* smallestInfiniteSet.addBack(2);    // 2 已经在集合中，所以不做任何变更。
* smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 是最小的整数，并将其从集合中移除。
* smallestInfiniteSet.popSmallest(); // 返回 2 ，并将其从集合中移除。
* smallestInfiniteSet.popSmallest(); // 返回 3 ，并将其从集合中移除。
* smallestInfiniteSet.addBack(1);    // 将 1 添加到该集合中。
* smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 在上一步中被添加到集合中，
* ⁠                                  // 且 1 是最小的整数，并将其从集合中移除。
* smallestInfiniteSet.popSmallest(); // 返回 4 ，并将其从集合中移除。
* smallestInfiniteSet.popSmallest(); // 返回 5 ，并将其从集合中移除。
*
*
*
* 提示：
*
*
* 1 <= num <= 1000
* 最多调用 popSmallest 和 addBack 方法 共计 1000 次
*
*
*/
package jzoffer

import (
	"container/heap"
	"sort"
)

// @lc code=start
type SmallestInfiniteSet struct {
	sort.IntSlice
	in map[int]bool
}

func (s *SmallestInfiniteSet) Push(v interface{}) {
	s.IntSlice = append(s.IntSlice, v.(int))
}

func (s *SmallestInfiniteSet) Pop() interface{} {
	if len(s.IntSlice) == 0 {
		return nil
	}
	top := s.IntSlice[len(s.IntSlice)-1]
	s.IntSlice = s.IntSlice[:len(s.IntSlice)-1]
	return top
}

func Constructor() SmallestInfiniteSet {
	nums := make([]int, 1000)
	m := make(map[int]bool)
	for i := 0; i < 1000; i++ {
		nums[i] = i + 1
		m[i+1] = true
	}
	res := &(SmallestInfiniteSet{IntSlice: nums, in: m})
	heap.Init(res)
	return *res
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	v := heap.Pop(this)
	if v == nil {
		return -1
	}
	delete(this.in, v.(int))
	return v.(int)
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.in[num]; ok {
		return
	}
	this.in[num] = true
	heap.Push(this, num)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
// @lc code=end
