/*
 * @lc app=leetcode.cn id=2347 lang=golang
 *
 * [2347] 最好的扑克手牌
 *
 * https://leetcode.cn/problems/best-poker-hand/description/
 *
 * algorithms
 * Easy (58.93%)
 * Likes:    23
 * Dislikes: 0
 * Total Accepted:    12.1K
 * Total Submissions: 20.5K
 * Testcase Example:  '[13,2,3,1,9]\n["a","a","a","a","a"]'
 *
 * 给你一个整数数组 ranks 和一个字符数组 suit 。你有 5 张扑克牌，第 i 张牌大小为 ranks[i] ，花色为 suits[i] 。
 *
 * 下述是从好到坏你可能持有的 手牌类型 ：
 *
 *
 * "Flush"：同花，五张相同花色的扑克牌。
 * "Three of a Kind"：三条，有 3 张大小相同的扑克牌。
 * "Pair"：对子，两张大小一样的扑克牌。
 * "High Card"：高牌，五张大小互不相同的扑克牌。
 *
 *
 * 请你返回一个字符串，表示给定的 5 张牌中，你能组成的 最好手牌类型 。
 *
 * 注意：返回的字符串 大小写 需与题目描述相同。
 *
 *
 *
 * 示例 1：
 *
 * 输入：ranks = [13,2,3,1,9], suits = ["a","a","a","a","a"]
 * 输出："Flush"
 * 解释：5 张扑克牌的花色相同，所以返回 "Flush" 。
 *
 *
 * 示例 2：
 *
 * 输入：ranks = [4,4,2,4,4], suits = ["d","a","a","b","c"]
 * 输出："Three of a Kind"
 * 解释：第一、二和四张牌组成三张相同大小的扑克牌，所以得到 "Three of a Kind" 。
 * 注意我们也可以得到 "Pair" ，但是 "Three of a Kind" 是更好的手牌类型。
 * 有其他的 3 张牌也可以组成 "Three of a Kind" 手牌类型。
 *
 * 示例 3：
 *
 * 输入：ranks = [10,10,2,12,9], suits = ["a","b","c","a","d"]
 * 输出："Pair"
 * 解释：第一和第二张牌大小相同，所以得到 "Pair" 。
 * 我们无法得到 "Flush" 或者 "Three of a Kind" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * ranks.length == suits.length == 5
 * 1 <= ranks[i] <= 13
 * 'a' <= suits[i] <= 'd'
 * 任意两张扑克牌不会同时有相同的大小和花色。
 *
 *
 */
package jzoffer

import "bytes"

// @lc code=start
func bestHand(ranks []int, suits []byte) (res string) {
	// "Flush"：同花，五张相同花色的扑克牌。
	// "Three of a Kind"：三条，有 3 张大小相同的扑克牌。
	// "Pair"：对子，两张大小一样的扑克牌。
	// "High Card"：高牌，五张大小互不相同的扑克牌
	pre, equal := suits[0], true
	rm := make(map[int]int)
	for i := 0; i < len(suits); i++ {
		if i > 0 && suits[i] != pre {
			equal = false
		}
		rm[ranks[i]]++
	}
	if equal {
		return "Flush"
	}
	n2, n3 := false, false
	for _, v := range rm {
		if v >= 3 {
			n3 = true
		}
		if v >= 2 {
			n2 = true
		}
	}
	if n3 {
		return "Three of a Kind"
	}
	if n2 {
		return "Pair"
	}
	// if len(rm) == 5 {
	// 	return "High Card"
	// }
	// return ""
	return "High Card"
}

func bestHand0(ranks []int, suits []byte) string {
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	cnt, pair := map[int]int{}, false
	for _, r := range ranks {
		cnt[r]++
		if cnt[r] == 3 {
			return "Three of a Kind"
		}
		if cnt[r] == 2 {
			pair = true
		}
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}

// @lc code=end
