/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode.cn/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (45.38%)
 * Likes:    1599
 * Dislikes: 0
 * Total Accepted:    250.7K
 * Total Submissions: 552.4K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 *
 *
 */
package jzoffer

// @lc code=start
func subarraySum(nums []int, k int) int {
	n := len(nums)
	var res, sum int
	for i := 0; i < n; i++ {
		sum = nums[i]
		if sum == k {
			res++
		}
		for j := i + 1; j < n; j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

// 前缀和加hash优化
// m[i]为前缀和为i的个数 j到i的子数组和为k: sum[i]-sum[j-1]==k => sum[i]-k==sum[j-1]
func subarraySumHash(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/QTMn0o/solution/he-wei-k-de-zi-shu-zu-by-leetcode-soluti-1169/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
