package jzoffer

// 给你一个下标从 0 开始的字符串数组 words 以及一个二维整数数组 queries 。
// 每个查询 queries[i] = [li, ri] 会要求我们统计在 words 中下标在 li 到 ri 范围内（包含 这两个值）并且以元音开头和结尾的字符串的数目。
// 返回一个整数数组，其中数组的第 i 个元素对应第 i 个查询的答案。
// 注意：元音字母是 'a'、'e'、'i'、'o' 和 'u' 。
//
// 示例 1：
// 输入：words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
// 输出：[2,3,0]
// 解释：以元音开头和结尾的字符串是 "aba"、"ece"、"aa" 和 "e" 。
// 查询 [0,2] 结果为 2（字符串 "aba" 和 "ece"）。
// 查询 [1,4] 结果为 3（字符串 "ece"、"aa"、"e"）。
// 查询 [1,1] 结果为 0 。
// 返回结果 [2,3,0] 。
// 示例 2：
// 输入：words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
// 输出：[3,2,1]
// 解释：每个字符串都满足这一条件，所以返回 [3,2,1] 。
//
// 提示：
// 1 <= words.length <= 105
// 1 <= words[i].length <= 40
// words[i] 仅由小写英文字母组成
// sum(words[i].length) <= 3 * 105
// 1 <= queries.length <= 105
// 0 <= queries[j][0] <= queries[j][1] < words.length
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/count-vowel-strings-in-ranges
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func vowelStrings(words []string, queries [][]int) []int {
	num := make([]int, len(words))
	sum := make([]int, len(words)+1)
	for i, v := range words {
		switch v[0] {
		case 'a', 'e', 'i', 'o', 'u':
		default:
			continue
		}
		switch v[len(v)-1] {
		case 'a', 'e', 'i', 'o', 'u':
		default:
			continue
		}
		num[i] = 1
	}
	for i := 1; i <= len(num); i++ {
		sum[i] = sum[i-1] + num[i-1]
	}
	res := make([]int, len(queries))
	for i, v := range queries {
		l, r := v[0], v[1]
		res[i] = sum[r+1] - sum[l] // 1 0 2
	}
	return res
}
