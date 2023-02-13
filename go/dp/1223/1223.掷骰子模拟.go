/*
 * @lc app=leetcode.cn id=1223 lang=golang
 *
 * [1223] 掷骰子模拟
 *
 * https://leetcode.cn/problems/dice-roll-simulation/description/
 *
 * algorithms
 * Hard (49.80%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    7.4K
 * Total Submissions: 13K
 * Testcase Example:  '2\n[1,1,2,2,2,3]'
 *
 * 有一个骰子模拟器会每次投掷的时候生成一个 1 到 6 的随机数。
 *
 * 不过我们在使用它时有个约束，就是使得投掷骰子时，连续 掷出数字 i 的次数不能超过 rollMax[i]（i 从 1 开始编号）。
 *
 * 现在，给你一个整数数组 rollMax 和一个整数 n，请你来计算掷 n 次骰子可得到的不同点数序列的数量。
 *
 * 假如两个序列中至少存在一个元素不同，就认为这两个序列是不同的。由于答案可能很大，所以请返回 模 10^9 + 7 之后的结果。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 2, rollMax = [1,1,2,2,2,3]
 * 输出：34
 * 解释：我们掷 2 次骰子，如果没有约束的话，共有 6 * 6 = 36 种可能的组合。但是根据 rollMax 数组，数字 1 和 2
 * 最多连续出现一次，所以不会出现序列 (1,1) 和 (2,2)。因此，最终答案是 36-2 = 34。
 *
 *
 * 示例 2：
 *
 * 输入：n = 2, rollMax = [1,1,1,1,1,1]
 * 输出：30
 *
 *
 * 示例 3：
 *
 * 输入：n = 3, rollMax = [1,1,1,2,2,3]
 * 输出：181
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 5000
 * rollMax.length == 6
 * 1 <= rollMax[i] <= 15
 *
 *
 */
package jzoffer

// @lc code=start

// TODO

// 由回溯推导dp解方程
// https://leetcode.cn/problems/dice-roll-simulation/solution/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-sje6/
// 会超时的回溯写法
func dieSimulatorBackTracing(n int, rollMax []int) (ans int) {
	const mod int = 1e9 + 7
	var dfs func(int, int, int) int
	// i:剩余掷骰子的次数
	// last: 上次的骰子数值
	// left,last的剩余出现次数
	dfs = func(i, last, left int) (res int) {
		if i == 0 {
			return 1
		}
		for j, mx := range rollMax {
			if j != last {
				res += dfs(i-1, j, mx-1)
			} else if left > 0 {
				res += dfs(i-1, j, left-1)
			}
		}
		return res % mod
	}
	for j, mx := range rollMax {
		ans += dfs(n-1, j, mx-1)
	}
	return ans % mod
}

// 1. 优化,减少重复搜索
func dieSimulatorOptimizedBacktracking(n int, rollMax []int) (ans int) {
	const mod int = 1e9 + 7
	const m = 6
	cache := make([][m][]int, n)
	for i := range cache {
		for j := range cache[i] {
			cache[i][j] = make([]int, rollMax[j])
			for k := range cache[i][j] {
				cache[i][j][k] = -1 // -1 表示没有访问过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, last, left int) (res int) {
		if i == 0 {
			return 1
		}
		c := &cache[i][last][left]
		if *c != -1 {
			return *c
		}
		for j, mx := range rollMax {
			if j != last {
				res += dfs(i-1, j, mx-1)
			} else if left > 0 {
				res += dfs(i-1, j, left-1)
			}
		}
		res %= mod
		*c = res
		return
	}
	for j, mx := range rollMax {
		ans += dfs(n-1, j, mx-1)
	}
	return ans % mod
}

// 2.翻译成dp
func dieSimulator(n int, rollMax []int) (ans int) {
	// DP(pos, last) which means we are at the position pos having as last the last character seen.
	const mod int = 1e9 + 7
	const m = 6
	f := make([][m][]int, n)
	for i := range f {
		for j := range f[i] {
			f[i][j] = make([]int, rollMax[j])
		}
	}
	for j := range f[0] {
		for k := range f[0][j] {
			f[0][j][k] = 1
		}
	}
	for i := 1; i < n; i++ { // 次数
		for last, mx0 := range rollMax { //上次的值,连续次数
			for left := 0; left < mx0; left++ {
				res := 0
				for j, mx := range rollMax {
					if j != last {
						res += f[i-1][j][mx-1]
					} else if left > 0 {
						res += f[i-1][j][left-1]
					}
				}
				f[i][last][left] = res % mod
			}
		}
	}
	for j, mx := range rollMax {
		ans += f[n-1][j][mx-1]
	}
	return ans % mod
}

// @lc code=end
