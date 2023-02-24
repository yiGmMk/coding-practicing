/*
 * @lc app=leetcode.cn id=2357 lang=golang
 *
 * [2357] 使数组中所有元素都等于零
 *
 * https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/description/
 *
 * algorithms
 * Easy (74.22%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    37.6K
 * Total Submissions: 49.3K
 * Testcase Example:  '[1,5,0,3,5]'
 *
 * 给你一个非负整数数组 nums 。在一步操作中，你必须：
 *
 *
 * 选出一个正整数 x ，x 需要小于或等于 nums 中 最小 的 非零 元素。
 * nums 中的每个正整数都减去 x。
 *
 *
 * 返回使 nums 中所有元素都等于 0 需要的 最少 操作数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,0,3,5]
 * 输出：3
 * 解释：
 * 第一步操作：选出 x = 1 ，之后 nums = [0,4,0,2,4] 。
 * 第二步操作：选出 x = 2 ，之后 nums = [0,2,0,0,2] 。
 * 第三步操作：选出 x = 2 ，之后 nums = [0,0,0,0,0] 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：0
 * 解释：nums 中的每个元素都已经是 0 ，所以不需要执行任何操作。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 100
 *
 *
 */
package jzoffer

import (
	"sort"
)

// @lc code=start
func minimumOperations(nums []int) (res int) {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		sum += (nums[i] - sum)
		for j := i + 1; j < len(nums) && nums[j] == sum; j++ {
			i = j
		}
		res++
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/solution/shi-shu-zu-zhong-suo-you-yuan-su-du-deng-ix12/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func minimumOperations1(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		if x > 0 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}

// @lc code=end
