/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode.cn/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.97%)
 * Likes:    1857
 * Dislikes: 0
 * Total Accepted:    955.8K
 * Total Submissions: 1.5M
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
// 提示: A two-pointer approach could be helpful here.
// The idea would be to have one pointer for iterating the array and another pointer that just works on the non-zero elements of the array.
func moveZeroes1(nums []int) {
	// p0找是0的数,p找非0的
	for p0, p := 0, 0; p0 < len(nums) && p < len(nums); p0++ {
		if nums[p0] != 0 {
			p++
			continue
		}
		for p < len(nums) && nums[p] == 0 {
			p++
		}
		if p0 < p && p < len(nums) { //只能把后面的非0的数与前面的0交换
			nums[p0], nums[p] = nums[p], nums[p0]
		}
	}
	return
}

// 参考 :https://leetcode.cn/problems/move-zeroes/solution/dong-hua-yan-shi-283yi-dong-ling-by-wang_ni_ma/
// 两次遍历
func moveZeroes2(nums []int) {
	var p int
	//第一次遍历的时候，p指针记录非0的个数，只要是非0的统统都赋给nums[p]
	for i, v := range nums {
		if v != 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	for i := p; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 一次遍历
// 这里参考了快速排序的思想，快速排序首先要确定一个待分割的元素做中间点x，然后把所有小于等于x的元素放到x的左边，大于x的元素放到其右边。
// 这里我们可以用0当做这个中间点，把不等于0(注意题目没说不能有负数)的放到中间点的左边，等于0的放到其右边。
// 这的中间点就是0本身，所以实现起来比快速排序简单很多，我们使用两个指针i和j，只要nums[i]!=0，我们就交换nums[i]和nums[j]
// 作者：wang_ni_ma
// 链接：https://leetcode.cn/problems/move-zeroes/solution/dong-hua-yan-shi-283yi-dong-ling-by-wang_ni_ma/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func moveZeroes(nums []int) {
	var j int
	for i, v := range nums {
		if v != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/move-zeroes/solution/yi-dong-ling-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func moveZeroes3(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// @lc code=end
