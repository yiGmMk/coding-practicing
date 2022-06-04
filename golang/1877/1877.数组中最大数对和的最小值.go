/*
 * @lc app=leetcode.cn id=1877 lang=golang
 *
 * [1877] 数组中最大数对和的最小值
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

// @lc code=start
func minPairSum(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sort.Ints(nums)
	fmt.Println(nums)
	max := math.MinInt64
	l := len(nums)
	for id := 0; id <= l/2; id++ {
		if sum := (nums[id] + nums[l-id-1]); sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end
