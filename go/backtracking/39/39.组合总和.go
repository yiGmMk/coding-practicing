/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode.cn/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.34%)
 * Likes:    2684
 * Dislikes: 0
 * Total Accepted:    829.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target
 * 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
 *
 * candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
 *
 * 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：candidates = [2,3,6,7], target = 7
 * 输出：[[2,2,3],[7]]
 * 解释：
 * 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
 * 7 也是一个候选， 7 = 7 。
 * 仅有这两种组合。
 *
 * 示例 2：
 *
 *
 * 输入: candidates = [2,3,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
 *
 * 示例 3：
 *
 *
 * 输入: candidates = [2], target = 1
 * 输出: []
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 2 <= candidates[i] <= 40
 * candidates 的所有元素 互不相同
 * 1 <= target <= 40
 *
 *
 */
package jzoffer

import (
	"fmt"
	"sort"
)

// @lc code=start
func combinationSumMy(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	visited := make(map[string]bool)
	var backtrace func(left int, path []int)
	backtrace = func(left int, path []int) {
		if left < 0 {
			return
		}
		if left == 0 {
			sort.Ints(path)
			if _, ok := visited[fmt.Sprintf("%v", path)]; ok {
				return
			}
			visited[fmt.Sprintf("%v", path)] = true
			// item := make([]int, len(path))
			// copy(item, path)
			res = append(res, path)
			return
		}
		for _, v := range candidates {
			if left-v < 0 {
				break
			}
			path = append(path, v)
			// 必须copy,否则会影响backtrace的递归
			p := make([]int, len(path))
			copy(p, path)
			backtrace(left-v, p)
			path = path[:len(path)-1]
		}
	}
	backtrace(target, []int{})
	return
}

// 作者：ukcuf
// 链接：https://leetcode.cn/problems/combination-sum/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	ret := [][]int{}
	for i, d := range candidates {
		if target-d < 0 {
			break
		} else if target-d == 0 {
			ret = append(ret, []int{d})
			continue
		}
		for _, v := range dfs(candidates[i:], target-d) {
			ret = append(ret, append([]int{d}, v...))
		}
	}
	return ret
}

// @lc code=end
