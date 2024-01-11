/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.17%)
 * Likes:    2223
 * Dislikes: 0
 * Total Accepted:    716.4K
 * Total Submissions: 882.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */
package jzoffer

import (
	"fmt"
	"sort"
)

// @lc code=start
// 回溯
// 回溯
func subsetsV0(nums []int) (res [][]int) {
	//
	sort.Ints(nums)
	var dfs func(path []int, start int)
	dfs = func(path []int, start int) {
		if len(nums) == start {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(nums); i++ {
			// 子集不含
			dfs(path, i+1)
			// 子集包含 i
			dfs(append(path, nums[i]), i+1)
		}
		return
	}
	dfs([]int{}, 0)

	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})

	m := make(map[string]bool)
	out := [][]int{}
	for i := 0; i < len(res); i++ {
		str := fmt.Sprintf("%v", res[i])
		if _, ok := m[str]; !ok {
			out = append(out, res[i])
			m[str] = true
		}
	}
	return out
}

func subsets(nums []int) (res [][]int) {
	//
	sort.Ints(nums)
	var dfs func(path []int, start int)
	dfs = func(path []int, start int) {
		if len(nums) == start {
			res = append(res, append([]int{}, path...))
			return
		}
		// 子集不含
		dfs(path, start+1)
		// 子集包含 i
		dfs(append(path, nums[start]), start+1)

		return
	}
	dfs([]int{}, 0)
	return res
}

// @lc code=end
