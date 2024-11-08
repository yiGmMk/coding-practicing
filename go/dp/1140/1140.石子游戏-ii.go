/*
 * @lc app=leetcode.cn id=1140 lang=golang
 *
 * [1140] 石子游戏 II
 *
 * https://leetcode.cn/problems/stone-game-ii/description/
 *
 * algorithms
 * Medium (66.27%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    16.8K
 * Total Submissions: 24.2K
 * Testcase Example:  '[2,7,9,4,4]'
 *
 * 爱丽丝和鲍勃继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
 *
 * 爱丽丝和鲍勃轮流进行，爱丽丝先开始。最初，M = 1。
 *
 * 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
 *
 * 游戏一直持续到所有石子都被拿走。
 *
 * 假设爱丽丝和鲍勃都发挥出最佳水平，返回爱丽丝可以得到的最大数量的石头。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：piles = [2,7,9,4,4]
 * 输出：10
 * 解释：如果一开始Alice取了一堆，Bob取了两堆，然后Alice再取两堆。爱丽丝可以得到2 + 4 + 4 =
 * 10堆。如果Alice一开始拿走了两堆，那么Bob可以拿走剩下的三堆。在这种情况下，Alice得到2 + 7 = 9堆。返回10，因为它更大。
 *
 *
 * 示例 2:
 *
 *
 * 输入：piles = [1,2,3,4,5,100]
 * 输出：104
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= piles.length <= 100
 * 1 <= piles[i] <= 10^4
 *
 *
 */
package jzoffer

import "math"

// TODO:

// @lc code=start
// Use dynamic programming: the states are (i, m) for the answer of piles[i:]
// and that given m.
// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/stone-game-ii/solution/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-jjax/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func stoneGameIIDfs(s []int) int {
	n := len(s)
	for i := n - 2; i >= 0; i-- {
		s[i] += s[i+1] // 后缀和
	}

	cache := make([][]int, n-1)
	for i := range cache {
		cache[i] = make([]int, (n+1)/4+1)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没有访问过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, m int) int {
		if i+m*2 >= n {
			return s[i]
		}
		v := &cache[i][m]
		if *v != -1 {
			return *v
		}
		mn := math.MaxInt
		for x := 1; x <= m*2; x++ {
			mn = min(mn, dfs(i+x, max(m, x)))
		}
		// 对于每个节点，由于剩余的石子总数是固定的，如果拿了某几堆石子后，
		// 对手能得到的石子数最少，那么自己能得到的石子数就是最多的
		res := s[i] - mn
		*v = res
		return res
	}
	return dfs(0, 1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 等比数列求和公式: a1+a2+...an= (an*q -a1) / (q-1)
// m 最大为[i/2]+1 => 极端情况下每次拿2M堆,i=2+4+8+...+M= ( M*2 -2 ) /( 2-1 )=2M-2 ==> m<=i/2+1
func stoneGameII(piles []int) int {
	n := len(piles)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i, s := n-1, 0; i >= 0; i-- {
		s += piles[i]
		for m := 1; m <= i/2+1; m++ {
			if i+m*2 >= n {
				f[i][m] = s
			} else {
				// 对于每个节点，由于剩余的石子总数是固定的，如果拿了某几堆石子后，
				// 对手能得到的石子数最少，那么自己能得到的石子数就是最多的
				mn := math.MaxInt
				for x := 1; x <= m*2; x++ {
					mn = min(mn, f[i+x][max(m, x)])
				}
				f[i][m] = s - mn
			}
		}
	}
	return f[0][1]
}

// @lc code=end
