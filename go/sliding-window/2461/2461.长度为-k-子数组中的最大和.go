/*
 * @lc app=leetcode.cn id=2461 lang=golang
 *
 * [2461] 长度为 K 子数组中的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/
 *
 * algorithms
 * Medium (30.13%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    7.4K
 * Total Submissions: 24.5K
 * Testcase Example:  '[1,5,4,2,9,9,9]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k 。请你从 nums 中满足下述条件的全部子数组中找出最大子数组和：
 *
 *
 * 子数组的长度是 k，且
 * 子数组中的所有元素 各不相同 。
 *
 *
 * 返回满足题面要求的最大子数组和。如果不存在子数组满足这些条件，返回 0 。
 *
 * 子数组 是数组中一段连续非空的元素序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,5,4,2,9,9,9], k = 3
 * 输出：15
 * 解释：nums 中长度为 3 的子数组是：
 * - [1,5,4] 满足全部条件，和为 10 。
 * - [5,4,2] 满足全部条件，和为 11 。
 * - [4,2,9] 满足全部条件，和为 15 。
 * - [2,9,9] 不满足全部条件，因为元素 9 出现重复。
 * - [9,9,9] 不满足全部条件，因为元素 9 出现重复。
 * 因为 15 是满足全部条件的所有子数组中的最大子数组和，所以返回 15 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [4,4,4], k = 3
 * 输出：0
 * 解释：nums 中长度为 3 的子数组是：
 * - [4,4,4] 不满足全部条件，因为元素 4 出现重复。
 * 因为不存在满足全部条件的子数组，所以返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 *
 *
 */
package jzoffer

// @lc code=start
func maximumSubarraySum(nums []int, k int) int64 {
	m := make(map[int]int)
	res, sum := 0, 0

	l, r := 0, k-1
	for i := l; i < k-1; i++ {
		m[nums[i]]++
		sum += nums[i]
	}
	for ; r < len(nums); r++ {
		m[nums[r]]++
		sum += nums[r]
		if len(m) == k {
			res = max(res, sum)
		}
		if m[nums[l]] == 1 { // 有多个的不能直接删除
			delete(m, nums[l])
		} else {
			m[nums[l]]--
		}
		sum -= nums[l]
		l++
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/solution/hua-dong-chuang-kou-by-endlesscheng-m0gm/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func maximumSubarraySum1(nums []int, k int) int64 {
	ans, sum := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		sum += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++ // 移入元素
		sum += nums[i]
		if len(cnt) == k && sum > ans {
			ans = sum
		}
		x := nums[i+1-k]
		cnt[x]-- // 移出元素
		if cnt[x] == 0 {
			delete(cnt, x) // 重要：及时移除个数为 0 的数据
		}
		sum -= x
	}
	return int64(ans)
}

// @lc code=end
