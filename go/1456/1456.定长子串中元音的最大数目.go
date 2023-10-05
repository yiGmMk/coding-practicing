/*
 * @lc app=leetcode.cn id=1456 lang=golang
 *
 * [1456] 定长子串中元音的最大数目
 *
 * https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
 *
 * algorithms
 * Medium (55.89%)
 * Likes:    80
 * Dislikes: 0
 * Total Accepted:    38.3K
 * Total Submissions: 68.4K
 * Testcase Example:  '"abciiidef"\n3'
 *
 * 给你字符串 s 和整数 k 。
 *
 * 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
 *
 * 英文中的 元音字母 为（a, e, i, o, u）。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "abciiidef", k = 3
 * 输出：3
 * 解释：子字符串 "iii" 包含 3 个元音字母。
 *
 *
 * 示例 2：
 *
 * 输入：s = "aeiou", k = 2
 * 输出：2
 * 解释：任意长度为 2 的子字符串都包含 2 个元音字母。
 *
 *
 * 示例 3：
 *
 * 输入：s = "leetcode", k = 3
 * 输出：2
 * 解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
 *
 *
 * 示例 4：
 *
 * 输入：s = "rhythms", k = 4
 * 输出：0
 * 解释：字符串 s 中不含任何元音字母。
 *
 *
 * 示例 5：
 *
 * 输入：s = "tryhard", k = 4
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s 由小写英文字母组成
 * 1 <= k <= s.length
 *
 *
 */
package jzoffer

// @lc code=start
func maxVowelsMy(s string, k int) (res int) {
	// 比下面多了O(n)的存储空间
	size := make([]int, len(s))
	if is(s, 0) > 0 {
		size[0]++
		res = max(res, size[0])
	}

	for i := 1; i < len(s); i++ {
		size[i] += is(s, i) + size[i-1]
		if i >= k-1 {
			sum := size[i]
			if i > k-1 {
				sum -= size[i-k]
			}
			res = max(sum, res)
		}
	}
	return
}

func is(s string, i int) int {
	switch s[i] {
	case 'a', 'e', 'i', 'o', 'u':
		return 1

	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxVowels(s string, k int) (res int) {
	// k个里有多少元音
	kSize := 0
	for i := 0; i < k; i++ {
		kSize += is(s, i)
	}
	// 滑动窗口
	res = kSize
	for i := k; i < len(s); i++ {
		kSize += is(s, i) - is(s, i-k)
		res = max(res, kSize)
	}
	return
}

// @lc code=end
