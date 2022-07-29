/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (35.89%)
 * Likes:    5040
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */

package jzoffer

import (
	"sort"
	"strconv"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	index := -1 //最靠近零
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < 0 && nums[i+1] >= 0 {
			index = i
			break
		}
	}
	if index == -1 {
		if len(nums) >= 3 && nums[0]+nums[1]+nums[2] == 0 {
			res = append(res, []int{nums[0], nums[1], nums[2]})
		}
		return res
	}

	// 去重
	resMap := make(map[string]bool)

	for i := 0; i <= index+1 /*0不能漏掉*/ && i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[j]+nums[k] == target {
				strRes := strconv.Itoa(nums[i]) + "-" + strconv.Itoa(nums[j]) + "-" + strconv.Itoa(nums[k])
				if _, ok := resMap[strRes]; !ok {
					resMap[strRes] = true
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				j++
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// @lc code=end

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/1fGaJU/solution/shu-zu-zhong-he-wei-0-de-san-ge-shu-by-l-hzw9/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
