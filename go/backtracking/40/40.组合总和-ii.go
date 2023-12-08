/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (59.51%)
 * Likes:    1492
 * Dislikes: 0
 * Total Accepted:    481.3K
 * Total Submissions: 808.9K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	var backtrace func(nums []int, left int, path []int)
	backtrace = func(nums []int, left int, path []int) {
		if left == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if left < 0 {
			return
		}

		for i, v := range nums {
			if left-v < 0 {
				break
			}
			if i > 0 && v == nums[i-1] {
				continue
			}
			path = append(path, v)
			backtrace(nums[i+1:], left-v, append([]int{}, path...))
			path = path[:len(path)-1]
		}
	}

	backtrace(candidates, target, []int{})
	return
}

// @lc code=end
