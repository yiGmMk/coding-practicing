package jzoffer

import "sort"

/*
沿街有一排连续的房屋。每间房屋内都藏有一定的现金。现在有一位小偷计划从这些房屋中窃取现金。
由于相邻的房屋装有相互连通的防盗系统，所以小偷 不会窃取相邻的房屋 。
小偷的 窃取能力 定义为他在窃取过程中能从单间房屋中窃取的 最大金额 。
给你一个整数数组 nums 表示每间房屋存放的现金金额。形式上，从左起第 i 间房屋中放有 nums[i] 美元。
另给你一个整数数组 k ，表示窃贼将会窃取的 最少 房屋数。小偷总能窃取至少 k 间房屋。
返回小偷的 最小 窃取能力。
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/house-robber-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

提示：
1 <= nums.length <= 105
1 <= nums[i] <= 109
1 <= k <= (nums.length + 1)/2
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TODO
// https://leetcode.cn/problems/house-robber-iv/solution/er-fen-da-an-dp-by-endlesscheng-m558/
func minCapability1(nums []int, k int) int {
	check := func(mx int) bool {
		f0, f1 := 0, 0
		for _, v := range nums {
			if v < mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	}
	return sort.Search(1e9, check)
}

func minCapability(nums []int, k int) int {
	check := func(mx int) bool {
		dp := make([]int, len(nums))
		if nums[0] <= mx {
			dp[0] = 1
		}
		if len(nums) > 1 && (nums[0] <= mx || nums[1] <= mx) {
			dp[1] = 1
		}
		for i := 2; i < len(nums); i++ {
			if nums[i] <= mx {
				dp[i] = max(dp[i-1], dp[i-2]+1)
			} else {
				dp[i] = dp[i-1]
			}
		}
		return dp[len(nums)-1] >= k
	}
	return sort.Search(1e9, check)
}
