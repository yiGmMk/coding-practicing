/*
 * @lc app=leetcode.cn id=2289 lang=golang
 *
 * [2289] 使数组按非递减顺序排列
 *
 * https://leetcode.cn/problems/steps-to-make-array-non-decreasing/description/
 *
 * algorithms
 * Medium (22.00%)
 * Likes:    132
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 28.9K
 * Testcase Example:  '[5,3,4,4,7,3,6,11,8,5,11]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。在一步操作中，移除所有满足 nums[i - 1] > nums[i] 的 nums[i] ，其中 0
 * < i < nums.length 。
 *
 * 重复执行步骤，直到 nums 变为 非递减 数组，返回所需执行的操作数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,3,4,4,7,3,6,11,8,5,11]
 * 输出：3
 * 解释：执行下述几个步骤：
 * - 步骤 1 ：[5,3,4,4,7,3,6,11,8,5,11] 变为 [5,4,4,7,6,11,11]
 * - 步骤 2 ：[5,4,4,7,6,11,11] 变为 [5,4,7,11,11]
 * - 步骤 3 ：[5,4,7,11,11] 变为 [5,7,11,11]
 * [5,7,11,11] 是一个非递减数组，因此，返回 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,7,7,13]
 * 输出：0
 * 解释：nums 已经是一个非递减数组，因此，返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 *
 *
 */
package jzoffer

// @lc code=start
func totalSteps(nums []int) (ans int) {
	type pair struct{ v, t int }
	st := []pair{}
	for _, num := range nums {
		maxT := 0
		for len(st) > 0 && st[len(st)-1].v <= num {
			maxT = max(maxT, st[len(st)-1].t)
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			maxT++
			ans = max(ans, maxT)
		} else {
			maxT = 0
		}
		st = append(st, pair{num, maxT})
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/steps-to-make-array-non-decreasing/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
