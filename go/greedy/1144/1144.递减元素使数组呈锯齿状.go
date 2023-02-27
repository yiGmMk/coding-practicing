/*
 * @lc app=leetcode.cn id=1144 lang=golang
 *
 * [1144] 递减元素使数组呈锯齿状
 *
 * https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag/description/
 *
 * algorithms
 * Medium (44.97%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    14.2K
 * Total Submissions: 30.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums，每次 操作 会从中选择一个元素并 将该元素的值减少 1。
 *
 * 如果符合下列情况之一，则数组 A 就是 锯齿数组：
 *
 *
 * 每个偶数索引对应的元素都大于相邻的元素，即 A[0] > A[1] < A[2] > A[3] < A[4] > ...
 * 或者，每个奇数索引对应的元素都大于相邻的元素，即 A[0] < A[1] > A[2] < A[3] > A[4] < ...
 *
 *
 * 返回将数组 nums 转换为锯齿数组所需的最小操作次数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,2,3]
 * 输出：2
 * 解释：我们可以把 2 递减到 0，或把 3 递减到 1。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [9,6,1,6,2]
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 1000
 *
 *
 */
package jzoffer

// @lc code=start
func movesToMakeZigzag0(nums []int) (res int) {
	if len(nums) <= 2 {
		return 0
	}
	vs := make([]int, len(nums))
	copy(vs, nums)

	val := func(i int) (int, int) {
		v := nums[i-1]
		if i+1 >= len(nums) {
			return v, nums[i]
		}
		return v, nums[i+1]
	}
	r1, r2 := 0, 0
	for i := 1; i < len(nums); i += 2 {
		v, v1 := val(i)
		if v <= nums[i] {
			r1 += abs(v-nums[i]) + 1 // 偶数小
			nums[i] = v - 1
		}
		if v1 <= nums[i] && i+1 < len(nums) {
			r1 += abs(v1-nums[i]) + 1 // 偶数小
			nums[i] = v1 - 1
		}
	}
	nums = vs
	for i := 1; i < len(nums); i += 2 {
		v, v1 := val(i)
		if v >= nums[i] {
			r2 += abs(v-nums[i]) + 1 //偶数大
		}
		if v1 >= nums[i] && i+1 < len(nums) {
			r2 += abs(v1-nums[i]) + 1
			nums[i+1] = nums[i] - 1
		}
	}
	return min(r1, r2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

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

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag/solution/di-jian-yuan-su-shi-shu-zu-cheng-ju-chi-o30ye/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func movesToMakeZigzag(nums []int) int {
	help := func(pos int) int {
		res := 0
		for i := pos; i < len(nums); i += 2 {
			a := 0
			if i-1 >= 0 {
				a = max(a, nums[i]-nums[i-1]+1)
			}
			if i+1 < len(nums) {
				a = max(a, nums[i]-nums[i+1]+1)
			}
			res += a
		}
		return res
	}

	return min(help(0), help(1))
}

// @lc code=end
