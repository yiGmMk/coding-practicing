/*
 * @lc app=leetcode.cn id=1283 lang=golang
 *
 * [1283] 使结果不超过阈值的最小除数
 *
 * https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/description/
 *
 * algorithms
 * Medium (47.64%)
 * Likes:    83
 * Dislikes: 0
 * Total Accepted:    12.7K
 * Total Submissions: 26.7K
 * Testcase Example:  '[1,2,5,9]\n6'
 *
 * 给你一个整数数组 nums 和一个正整数 threshold  ，你需要选择一个正整数作为除数，然后将数组里每个数都除以它，并对除法结果求和。
 *
 * 请你找出能够使上述结果小于等于阈值 threshold 的除数中 最小 的那个。
 *
 * 每个数除以除数后都向上取整，比方说 7/3 = 3 ， 10/2 = 5 。
 *
 * 题目保证一定有解。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,5,9], threshold = 6
 * 输出：5
 * 解释：如果除数为 1 ，我们可以得到和为 17 （1+2+5+9）。
 * 如果除数为 4 ，我们可以得到和为 7 (1+1+2+3) 。如果除数为 5 ，和为 5 (1+1+1+2)。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,3,5,7,11], threshold = 11
 * 输出：3
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [19], threshold = 5
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * 1 <= nums[i] <= 10^6
 * nums.length <= threshold <= 10^6
 *
 *
 */
package jzoffer

import (
	"sort"
)

// TODO,类似287
// @lc code=start
func smallestDivisor1(nums []int, threshold int) (res int) {
	sort.Ints(nums)
	l, r := 0, nums[len(nums)-1]
	for l+1 < r {
		mid := int(uint(l+r) >> 1)
		sum := 0
		for _, v := range nums {
			//val := int(math.Ceil(float64(v / mid)))           // 向上取整,出错了
			//val1 := int(math.Ceil(float64(v) / float64(mid))) //
			val2 := (v-1)/mid + 1
			// if val != val2 {
			// 	fmt.Printf("v=%d,mid=%d,ceil=%d,ceil1=%d,ceil2=%d\n", v, mid, val, val1, val2)
			// }
			sum += val2
		}
		if sum > threshold {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func smallestDivisor(nums []int, threshold int) (res int) {
	sort.Ints(nums)
	l, r := 1, nums[len(nums)-1]
	for l <= r {
		mid := int(uint(l+r) >> 1)
		sum := 0
		for _, v := range nums {
			val2 := (v-1)/mid + 1
			sum += val2
		}
		if sum > threshold {
			l = mid + 1
		} else {
			r = mid - 1
			res = mid
		}
	}
	return res
}

// @lc code=end
