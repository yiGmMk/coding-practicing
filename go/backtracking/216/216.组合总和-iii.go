/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 *
 * https://leetcode.cn/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (71.17%)
 * Likes:    788
 * Dislikes: 0
 * Total Accepted:    320.3K
 * Total Submissions: 450.1K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
 *
 *
 * 只使用数字1到9
 * 每个数字 最多使用一次
 *
 *
 * 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 * 解释:
 * 1 + 2 + 4 = 7
 * 没有其他符合的组合了。
 *
 * 示例 2:
 *
 *
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 * 解释:
 * 1 + 2 + 6 = 9
 * 1 + 3 + 5 = 9
 * 2 + 3 + 4 = 9
 * 没有其他符合的组合了。
 *
 * 示例 3:
 *
 *
 * 输入: k = 4, n = 1
 * 输出: []
 * 解释: 不存在有效的组合。
 * 在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= k <= 9
 * 1 <= n <= 60
 *
 *
 */
package jzoffer

// @lc code=start
func combinationSum3(k int, n int) (res [][]int) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var dfs func(target int, nums, path []int)
	dfs = func(target int, nums, path []int) {
		if len(path) > k {
			return
		}
		if len(path) == k {
			if target == 0 {
				res = append(res, path)
			}
			return
		}
		for i := 0; i < len(nums); i++ {
			if target-nums[i] < 0 {
				break
			}
			path = append(path, nums[i])
			pick := append([]int{}, path...)
			dfs(target-nums[i], nums[i+1:], pick)
			path = path[:len(path)-1]
		}
	}
	dfs(n, nums, []int{})
	return
}

// TODO 二进制枚举, 最多9个数,不允许重复,每个数只有2中状态=> 那么最多2^9种状态
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/combination-sum-iii/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func combinationSum3Binary(k int, n int) (ans [][]int) {
	var temp []int
	check := func(mask int) bool {
		temp = nil
		sum := 0
		for i := 0; i < 9; i++ {
			if 1<<i&mask > 0 {
				temp = append(temp, i+1)
				sum += i + 1
			}
		}
		return len(temp) == k && sum == n
	}

	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			ans = append(ans, append([]int(nil), temp...))
		}
	}
	return
}

// @lc code=end
