/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 *
 * https://leetcode.cn/problems/kth-missing-positive-number/description/
 *
 * algorithms
 * Easy (54.04%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    36.8K
 * Total Submissions: 68.1K
 * Testcase Example:  '[2,3,4,7,11]\n5'
 *
 * 给你一个 严格升序排列 的正整数数组 arr 和一个整数 k 。
 *
 * 请你找到这个数组里第 k 个缺失的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [2,3,4,7,11], k = 5
 * 输出：9
 * 解释：缺失的正整数包括 [1,5,6,8,9,10,12,13,...] 。第 5 个缺失的正整数为 9 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,2,3,4], k = 2
 * 输出：6
 * 解释：缺失的正整数包括 [5,6,7,...] 。第 2 个缺失的正整数为 6 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 1000
 * 1 <= arr[i] <= 1000
 * 1 <= k <= 1000
 * 对于所有 1 <= i < j <= arr.length 的 i 和 j 满足 arr[i] < arr[j]
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以设计一个时间复杂度小于 O(n) 的算法解决此问题吗？
 *
 */
package jzoffer

import "sort"

// TODO redo

// @lc code=start
// !太捞了,没总结出规律(arr[i]前缺的数量可以计算出来)
func findKthPositiveOk(arr []int, k int) int {
	miss := []int{}
	for i, v := range arr {
		pre := 0
		if i > 0 {
			pre = arr[i-1]
		}
		if v-pre == 1 {
			continue
		}
		for j := pre + 1; j < v; j++ {
			miss = append(miss, j)
			if len(miss) == k {
				return j
			}
		}
	}
	if len(miss) < k {
		v := arr[len(arr)-1]
		for i := k - len(miss); i > 0; i-- {
			v++
		}
		return v
	}
	return -1
}

// 顺序枚举
// https://leetcode.cn/problems/kth-missing-positive-number/solution/di-k-ge-que-shi-de-zheng-zheng-shu-by-leetcode-sol/
func findKthPositiveOn1(arr []int, k int) int {
	missCount, lastMiss, current, ptr := 0, -1, 1, 0
	// current标记当前应该出现的数,
	// prt指向arr
	// missCount标记缺的数量
	for ; missCount < k; current++ {
		if current == arr[ptr] {
			if ptr+1 < len(arr) {
				ptr = ptr + 1
			}
		} else {
			missCount++
			lastMiss = current
		}
	}
	return lastMiss
}

// 二分+前缀和
// 存在的数有序,
func findKthPositive(arr []int, k int) int {
	missed := make([]int, len(arr))
	sum := make([]int, len(arr))
	for i, v := range arr {
		pre := 0
		if i > 0 {
			pre = arr[i-1]
		}
		if v-pre == 1 {
			continue
		}
		missed[i] = v - pre - 1
		sum[i] = missed[i]
	}
	for i := 1; i < len(sum); i++ {
		sum[i] += sum[i-1]
	}
	n := sort.SearchInts(sum, k)
	if n > 0 || n == len(arr) {
		return arr[n-1] + k - sum[n-1] //  需要减去中间缺的
	}
	return k //n为0说明缺的数据在最前面,k就是缺的数
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// 二分
// 对于每个元素 a[i],我们都可以唯一确定到第 i个元素为止缺失的元素数量为a[i]-i-1
func findKthPositive1(arr []int, k int) int {
	if arr[0] > k {
		return k
	}
	n := sort.Search(len(arr), func(i int) bool {
		return arr[i]-i-1 >= k
	})
	return arr[n-1] + k - (arr[n-1] - (n - 1) - 1)
}

// @lc code=end
