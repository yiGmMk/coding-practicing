/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.89%)
 * Likes:    2394
 * Dislikes: 0
 * Total Accepted:    794.9K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */

package jzoffer

// @lc code=start
func permute(nums []int) (res [][]int) {
	visited := make(map[int]bool)

	path := []int{}
	dfs(nums, path, &res, visited)
	return
}

// res为什么用指针类型?
// res必须传*[][]int,不然和外面传入的res空间分离了,返回空的[][]int
func dfs(nums []int, path []int, res *[][]int, visited map[int]bool) {
	if len(path) == len(nums) {
		// 思考下 *res = append(*res, path) 行不行
		//直接 append path会导致数据覆盖,path复用底层空间
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		path = append(path, nums[i])
		visited[nums[i]] = true
		dfs(nums, path, res, visited)
		visited[nums[i]] = false
		path = path[:len(path)-1]
	}
}

// @lc code=end
