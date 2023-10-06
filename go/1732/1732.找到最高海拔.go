/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] 找到最高海拔
 *
 * https://leetcode.cn/problems/find-the-highest-altitude/description/
 *
 * algorithms
 * Easy (81.35%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    55.9K
 * Total Submissions: 68.7K
 * Testcase Example:  '[-5,1,5,0,-7]'
 *
 * 有一个自行车手打算进行一场公路骑行，这条路线总共由 n + 1 个不同海拔的点组成。自行车手从海拔为 0 的点 0 开始骑行。
 *
 * 给你一个长度为 n 的整数数组 gain ，其中 gain[i] 是点 i 和点 i + 1 的 净海拔高度差（0 ）。请你返回 最高点的海拔
 * 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：gain = [-5,1,5,0,-7]
 * 输出：1
 * 解释：海拔高度依次为 [0,-5,-4,1,1,-6] 。最高海拔为 1 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：gain = [-4,-3,-2,-1,4,3,2]
 * 输出：0
 * 解释：海拔高度依次为 [0,-4,-7,-9,-10,-6,-3,-1] 。最高海拔为 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == gain.length
 * 1
 * -100
 *
 *
 */
package jzoffer

// @lc code=start
// func largestAltitude(gain []int) (res int) {

// h:=make([]int,len(gain)+1)
// res=math.MinInt
// for i:=1;i<=len(gain);i++{
//     h[i]=h[i-1]+gain[i-1]
// res=max(res,h[i])
// }
// return max(res,0)
// }

// func max(a,b int)int{
//     if a>b{
//         return a
//     }
//     return b
// }

func largestAltitude(gain []int) (ans int) {
	total := 0
	for _, x := range gain {
		total += x
		ans = max(ans, total)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
