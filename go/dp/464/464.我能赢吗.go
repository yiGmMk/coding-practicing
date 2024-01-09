/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] 我能赢吗
 *
 * https://leetcode.cn/problems/can-i-win/description/
 *
 * algorithms
 * Medium (40.80%)
 * Likes:    537
 * Dislikes: 0
 * Total Accepted:    40.5K
 * Total Submissions: 99.2K
 * Testcase Example:  '10\n11'
 *
 * 在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和 达到或超过  100
 * 的玩家，即为胜者。
 *
 * 如果我们将游戏规则改为 “玩家 不能 重复使用整数” 呢？
 *
 * 例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
 *
 * 给定两个整数 maxChoosableInteger （整数池中可选择的最大数）和 desiredTotal（累计和），若先出手的玩家能稳赢则返回
 * true ，否则返回 false 。假设两位玩家游戏时都表现 最佳 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：maxChoosableInteger = 10, desiredTotal = 11
 * 输出：false
 * 解释：
 * 无论第一个玩家选择哪个整数，他都会失败。
 * 第一个玩家可以选择从 1 到 10 的整数。
 * 如果第一个玩家选择 1，那么第二个玩家只能选择从 2 到 10 的整数。
 * 第二个玩家可以通过选择整数 10（那么累积和为 11 >= desiredTotal），从而取得胜利.
 * 同样地，第一个玩家选择任意其他整数，第二个玩家都会赢。
 *
 *
 * 示例 2:
 *
 *
 * 输入：maxChoosableInteger = 10, desiredTotal = 0
 * 输出：true
 *
 *
 * 示例 3:
 *
 *
 * 输入：maxChoosableInteger = 10, desiredTotal = 1
 * 输出：true
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= maxChoosableInteger <= 20
 * 0 <= desiredTotal <= 300
 *
 *
 */
package jzoffer

// @lc code=start

// TODO
// The premise for winning is pretty simple.
// If there exists a move given the current state that will grant the current player a win, then that state is a win, no matter what prior moves were made.
// On the contrary, if there are no possible moves so that the current player wins, then the player loses.
// Use a bitmask to represent the set of chosen numbers so far.
// Since the total number of possible bitmasks is low, use a slice instead of a map.
// Also, to avoid having to initialize the memoization slice, offset the player by one in the map, zero == not seen.
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	var sum int
	for val := 1; val <= maxChoosableInteger; val++ {
		sum += val
	}
	if sum < desiredTotal {
		return false
	}
	if desiredTotal <= maxChoosableInteger {
		return true
	}
	mem := make([]byte, (1<<(maxChoosableInteger+1))-1)
	res := findWinner(mem, 0, maxChoosableInteger, 0, desiredTotal, 0) == 0
	return res
}

// findWinner finds the winner. Returns 0 if the first player wins, otherwise 1.
// bm 记录已使用的数字
func findWinner(mem []byte, bm, max, sum, desiredTotal int, player byte) byte {
	if sum >= desiredTotal {
		return (player + 1) & 1
	}
	if mem[bm] != 0 {
		return mem[bm] - 1
	}
	// If any move is guaranteed to win, then this round is a win
	for i := 1; i <= max; i++ {
		if bm&(1<<i) > 0 {
			continue
		}
		winner := findWinner(mem, bm|(1<<i), max, sum+i, desiredTotal, (player+1)&1)
		if winner == player {
			mem[bm] = player + 1
			return player
		}
	}
	// No move could make this player win - return a loss
	mem[bm] = (player+1)&1 + 1
	return mem[bm] - 1
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/can-i-win/solutions/1506106/wo-neng-ying-ma-by-leetcode-solution-ef5v/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func canIWinDp(maxChoosableInteger, desiredTotal int) bool {
	// 等差数列求和 1 到 maxChoosableInteger
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(usedNum, curTot int) (res int8) {
		dv := &dp[usedNum]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := 0; i < maxChoosableInteger; i++ {
			if usedNum>>i&1 == 0 && (curTot+i+1 >= desiredTotal || dfs(usedNum|1<<i, curTot+i+1) == 0) {
				return 1
			}
		}
		return
	}
	return dfs(0, 0) == 1
}

// @lc code=end
