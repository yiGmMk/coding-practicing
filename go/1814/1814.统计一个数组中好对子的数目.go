/*
 * @lc app=leetcode.cn id=1814 lang=golang
 *
 * [1814] 统计一个数组中好对子的数目
 *
 * https://leetcode.cn/problems/count-nice-pairs-in-an-array/description/
 *
 * algorithms
 * Medium (36.99%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    20.8K
 * Total Submissions: 45.3K
 * Testcase Example:  '[42,11,1,97]'
 *
 * 给你一个数组 nums ，数组中只包含非负整数。定义 rev(x) 的值为将整数 x 各个数字位反转得到的结果。比方说 rev(123) = 321 ，
 * rev(120) = 21 。我们称满足下面条件的下标对 (i, j) 是 好的 ：
 *
 *
 * 0 <= i < j < nums.length
 * nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
 *
 *
 * 请你返回好下标对的数目。由于结果可能会很大，请将结果对 10^9 + 7 取余 后返回。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [42,11,1,97]
 * 输出：2
 * 解释：两个坐标对为：
 * ⁠- (0,3)：42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121 。
 * ⁠- (1,2)：11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [13,10,35,24,76]
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^9
 *
 *
 */
package joffer

// TODO redo

// @lc code=start
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/count-nice-pairs-in-an-array/solution/tong-ji-yi-ge-shu-zu-zhong-hao-dui-zi-de-ywux/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// func countNicePairs(nums []int) (ans int) {
// 	cnt := map[int]int{}
// 	for _, num := range nums {
// 		rev := 0
// 		for x := num; x > 0; x /= 10 {
// 			rev = rev*10 + x%10
// 		}
// 		ans += cnt[num-rev]
// 		cnt[num-rev]++
// 	}
// 	return ans % (1e9 + 7)
// }

// 1.The condition can be rearranged to
//
//	(nums[i] - rev(nums[i])) == (nums[j] - rev(nums[j]))
//
// 2.Transform each nums[i] into (nums[i] - rev(nums[i])).
// Then, count the number of (i, j) pairs that have equal values.
// 3. Keep a map storing the frequencies of values that you have seen so far.
// For each i, check if nums[i] is in the map. If it is,
// then add that count to the overall count.
// Then, increment the frequency of nums[i].
func countNicePairs(nums []int) (res int) {
	vm := make(map[int]int)
	rev := func(i int) int {
		out := 0
		for ; i > 0; i /= 10 {
			out = out*10 + i%10
		}
		return out
	}
	for _, v := range nums {
		r := v - rev(v)
		if num, ok := vm[r]; ok {
			res += num
		}
		vm[r]++
	}

	return res % (1e9 + 7)
}

// @lc code=end
