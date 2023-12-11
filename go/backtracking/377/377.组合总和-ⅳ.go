/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 *
 * https://leetcode.cn/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (52.74%)
 * Likes:    908
 * Dislikes: 0
 * Total Accepted:    166.7K
 * Total Submissions: 316K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
 *
 * 题目数据保证答案符合 32 位整数范围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3], target = 4
 * 输出：7
 * 解释：
 * 所有可能的组合为：
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * 请注意，顺序不同的序列被视作不同的组合。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [9], target = 3
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * nums 中的所有元素 互不相同
 * 1
 *
 *
 *
 *
 * 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
 *
 */
package jzoffer

// @lc code=start
// 直接用回溯会超时,使用map存储中间结果或者使用dp解答
func combinationSum4Backtrack(nums []int, target int) (res int) {
	mRes := make(map[int]int)
	var dfs func(cancidate []int, left int) int
	dfs = func(cancidate []int, left int) int {
		if left < 0 {
			return 0
		}
		if left == 0 {
			mRes[left] = 1
			return 1
		}
		num := 0
		for _, v := range nums {
			if left-v < 0 {
				continue
			}
			if n, ok := mRes[left-v]; ok {
				num += n
				continue
			}
			n := dfs(cancidate, left-v)
			mRes[left-v] = n
			num += n
		}
		return num
	}
	res += dfs(nums, target)
	return
}

func combinationSum4(nums []int, target int) (res int) {
	mRes := make(map[int]int)
	mRes[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i-v >= 0 {
				mRes[i] += mRes[i-v]
			}
		}
	}
	return mRes[target]
}

// @lc code=end
