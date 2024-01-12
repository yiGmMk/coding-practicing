/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode.cn/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (63.48%)
 * Likes:    1182
 * Dislikes: 0
 * Total Accepted:    341.8K
 * Total Submissions: 538.3K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
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
 *
 *
 *
 *
 */
package jzoffer

import (
	"fmt"
	"sort"
)

// @lc code=start

func subsetsWithDup(nums []int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				fmt.Printf("mask:%d.%b, idx:%d, v:%d\n", mask, mask, i, v)
				// 前一个数目和当前数相同，且前一个数之前没有出现过，
				// 说明是重复的，直接跳过
				// mask:1.1,   idx:0, v:1
				// mask:2.10,  idx:1, v:2
				// mask:3.11,  idx:0, v:1
				// mask:3.11,  idx:1, v:2
				// mask:4.100, idx:2, v:2
				// mask:5.101, idx:0, v:1
				// mask:5.101, idx:2, v:2
				// mask:6.110, idx:1, v:2
				// mask:6.110, idx:2, v:2
				// mask:7.111, idx:0, v:1
				// mask:7.111, idx:1, v:2
				// mask:7.111, idx:2, v:2
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		res = append(res, append([]int{}, t...))
	}
	return
}

func subsetsWithDupBT(nums []int) (res [][]int) {
	sort.Ints(nums)
	var dfs func(path []int, idx int)
	dfs = func(path []int, idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		dfs(path, idx+1)
		dfs(append(path, nums[idx]), idx+1)
		return
	}
	dfs([]int{}, 0)

	// 有重复的,去重
	m := make(map[string]bool, len(res))
	out := make([][]int, 0)
	for _, v := range res {
		str := fmt.Sprintf("%v", v)
		if !m[str] {
			m[str] = true
			out = append(out, v)
		}
	}
	return out
}

func subsetsWithDupBT2(nums []int) (res [][]int) {
	sort.Ints(nums)
	var dfs func(path []int, idx int)
	dfs = func(path []int, idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		dfs(append(path, nums[idx]), idx+1)

		// 不选就跳过后面一样的数，只需要用【78. 子集】的代码加这两行就搞定了！
		for i := idx + 1; i < len(nums); i++ {
			if nums[i] == nums[idx] {
				idx++
			}
		}
		dfs(path, idx+1)
		return
	}
	dfs([]int{}, 0)
	return
}

// @lc code=end
