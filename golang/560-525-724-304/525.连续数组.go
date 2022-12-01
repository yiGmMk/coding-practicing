/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 *
 * https://leetcode.cn/problems/contiguous-array/description/
 *
 * algorithms
 * Medium (54.47%)
 * Likes:    573
 * Dislikes: 0
 * Total Accepted:    58.7K
 * Total Submissions: 107.7K
 * Testcase Example:  '[0,1]'
 *
 * 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [0,1]
 * 输出: 2
 * 说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
 *
 * 示例 2:
 *
 *
 * 输入: nums = [0,1,0]
 * 输出: 2
 * 说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * nums[i] 不是 0 就是 1
 *
 *
 */
package jzoffer

// TODO
/* 由于「0 和 1 的数量相同」等价于「1 的数量减去 0 的数量等于 0」,
我们可以将数组中的 0 视作 -1,将问题转化为求和为0的最长连续数组 */
// @lc code=start
func findMaxLength(nums []int) (maxLength int) {
	mp := map[int]int{0: -1}
	counter := 0
	//sums := make([]int, len(nums)+1)
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		//sums[i] = counter
		if prevIndex, has := mp[counter]; has { // 前缀和相等,中间的一段前缀和是0
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	//fmt.Println(sums)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/A1NYOS/solution/0-he-1-ge-shu-xiang-tong-de-zi-shu-zu-by-xbyt/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
