/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 *
 * https://leetcode.cn/problems/can-place-flowers/description/
 *
 * algorithms
 * Easy (32.03%)
 * Likes:    625
 * Dislikes: 0
 * Total Accepted:    186.9K
 * Total Submissions: 583.9K
 * Testcase Example:  '[1,0,0,0,1]\n1'
 *
 * 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
 *
 * 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n
 * ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：flowerbed = [1,0,0,0,1], n = 1
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：flowerbed = [1,0,0,0,1], n = 2
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= flowerbed.length <= 2 * 10^4
 * flowerbed[i] 为 0 或 1
 * flowerbed 中不存在相邻的两朵花
 * 0 <= n <= flowerbed.length
 *
 */
package jzoffer

// @lc code=start
// 连续3个0种一朵花 https://leetcode.cn/problems/can-place-flowers/solutions/248589/lian-xu-san-ge-0ze-ke-yi-chong-yi-pen-hua-by-ya-le/?envType=study-plan-v2&envId=leetcode-75
func canPlaceFlowers(flowerbed []int, n int) bool {
	f := append([]int{0}, append(flowerbed, 0)...)
	for i := 1; i < len(f)-1; i++ {
		if f[i-1]+f[i]+f[i+1] == 0 {
			n--
			f[i] = 1
		}
	}

	return n <= 0
}

// @lc code=end
