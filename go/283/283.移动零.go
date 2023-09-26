/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode.cn/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.65%)
 * Likes:    2181
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 1.8M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [0]
 * 输出: [0]
 *
 *
 *
 * 提示:
 *
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能尽量减少完成的操作次数吗？
 *
 */
package jzoffer

// @lc code=start
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		}
		j := i + 1
		for ; j < len(nums) && nums[j] == 0; j++ {
		}
		if j >= len(nums) || nums[j] == 0 {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// func moveZeroes(nums []int) {
//     left, right, n := 0, 0, len(nums)
//     for right < n {
//         if nums[right] != 0 {
//             nums[left], nums[right] = nums[right], nums[left]
//             left++
//         }
//         right++
//     }
// }

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/move-zeroes/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
