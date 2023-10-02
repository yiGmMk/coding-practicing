/*
 * @lc app=leetcode.cn id=1679 lang=golang
 *
 * [1679] K 和数对的最大数目
 *
 * https://leetcode.cn/problems/max-number-of-k-sum-pairs/description/
 *
 * algorithms
 * Medium (55.60%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    23.4K
 * Total Submissions: 42K
 * Testcase Example:  '[1,2,3,4]\n5'
 *
 * 给你一个整数数组 nums 和一个整数 k 。
 *
 * 每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
 *
 * 返回你可以对数组执行的最大操作数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,4], k = 5
 * 输出：2
 * 解释：开始时 nums = [1,2,3,4]：
 * - 移出 1 和 4 ，之后 nums = [2,3]
 * - 移出 2 和 3 ，之后 nums = []
 * 不再有和为 5 的数对，因此最多执行 2 次操作。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,1,3,4,3], k = 6
 * 输出：1
 * 解释：开始时 nums = [3,1,3,4,3]：
 * - 移出前两个 3 ，之后nums = [1,4,3]
 * 不再有和为 6 的数对，因此最多执行 1 次操作。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
func maxOperations1(nums []int, k int) (res int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	// fmt.Println(m)
	for _, v := range nums {
		if m[v] > 0 && (m[k-v] > 0 && k-v != v) || (k-v == v && m[v] > 1) {
			m[k-v]--
			m[v]--
			res++
			//  fmt.Println(v,k-v)
		}
	}
	return res
}

func maxOperations(nums []int, k int) (res int) {
	sort.Ints(nums)

	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum > k {
			j--
			continue
		}
		if sum < k {
			i++
			continue
		}
		if sum == k {
			i++
			j--
			res++
		}
	}

	return res
}

// @lc code=end
