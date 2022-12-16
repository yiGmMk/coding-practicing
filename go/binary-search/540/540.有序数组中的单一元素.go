/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 *
 * https://leetcode.cn/problems/single-element-in-a-sorted-array/description/
 *
 * algorithms
 * Medium (60.60%)
 * Likes:    565
 * Dislikes: 0
 * Total Accepted:    108.7K
 * Total Submissions: 179.4K
 * Testcase Example:  '[1,1,2,3,3,4,4,8,8]'
 *
 * 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
 *
 * 请你找出并返回只出现一次的那个数。
 *
 * 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,2,3,3,4,4,8,8]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums =  [3,3,7,7,10,11,11]
 * 输出: 10
 *
 *
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^5
 *
 *
 */
package jzoffer

// @lc code=start
func singleNonDuplicate(nums []int) int {
	// i==i+1
	isBlue := func(i int) bool {
		// 偶数位置,i==i+1  => blue => true
		// 偶数位置,i!=i+1  => red => false
		if i == len(nums)-1 {
			return true
		}
		if i%2 == 0 {
			return nums[i] == nums[i+1]
		}
		return nums[i] == nums[i-1]
	}
	l, r := -1, len(nums)-1
	for l+1 < r {
		mid := int(uint(l+r) >> 1)
		if isBlue(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return nums[r]
}

func singleNonDuplicate3(nums []int) int {
	// i==i+1
	isBlue := func(i int) bool {
		// 偶数位置,i==i+1  => blue => true
		// 偶数位置,i!=i+1  => red => false
		if i == len(nums)-1 {
			return true
		}
		return nums[i] == nums[i^1]
	}
	l, r := -1, len(nums)-1
	for l+1 < r {
		mid := int(uint(l+r) >> 1)
		if isBlue(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return nums[r]
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/single-element-in-a-sorted-array/solution/you-xu-shu-zu-zhong-de-dan-yi-yuan-su-by-y8gh/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func singleNonDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func singleNonDuplicate2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}

// @lc code=end
