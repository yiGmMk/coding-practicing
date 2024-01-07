package jzoffer

// 给你两个下标从 0 开始的字符串 source 和 target ，它们的长度均为 n 并且由 小写 英文字母组成。
// 另给你两个下标从 0 开始的字符数组 original 和 changed ，以及一个整数数组 cost ，其中 cost[i] 代表将字符 original[i] 更改为字符 changed[i] 的成本。
// 你从字符串 source 开始。在一次操作中，如果 存在 任意 下标 j 满足 cost[j] == z  、original[j] == x 以及 changed[j] == y 。你就可以选择字符串中的一个字符 x 并以 z 的成本将其更改为字符 y 。
// 返回将字符串 source 转换为字符串 target 所需的 最小 成本。如果不可能完成转换，则返回 -1 。
// 注意，可能存在下标 i 、j 使得 original[j] == original[i] 且 changed[j] == changed[i] 。
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	dp := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < len(original); i++ {
		// 注意，可能存在下标 i 、j 使得 original[j] == original[i] 且 changed[j] == changed[i]
		if dp[original[i]-'a'][changed[i]-'a'] == -1 {
			dp[original[i]-'a'][changed[i]-'a'] = cost[i]
			continue
		}
		dp[original[i]-'a'][changed[i]-'a'] = min(dp[original[i]-'a'][changed[i]-'a'], cost[i])
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dp[i][k] != -1 && dp[k][j] != -1 {
					if dp[i][j] == -1 {
						dp[i][j] = dp[i][k] + dp[k][j]
					} else {
						dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]) // I,j 为-1的会更新失败
					}
				}
			}
		}
	}

	res := int64(0)
	for i := 0; i < len(source); i++ {
		if source[i] == target[i] {
			continue
		}
		if dp[source[i]-'a'][target[i]-'a'] == -1 {
			return -1
		}
		res += int64(dp[source[i]-'a'][target[i]-'a'])
	}
	return int64(res)
}
