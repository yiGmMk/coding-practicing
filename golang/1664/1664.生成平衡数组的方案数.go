/*
 * @lc app=leetcode.cn id=1664 lang=golang
 *
 * [1664] 生成平衡数组的方案数
 *
 * https://leetcode.cn/problems/ways-to-make-a-fair-array/description/
 *
 * algorithms
 * Medium (55.85%)
 * Likes:    51
 * Dislikes: 0
 * Total Accepted:    10.8K
 * Total Submissions: 17.6K
 * Testcase Example:  '[2,1,6,4]'
 *
 * 给你一个整数数组 nums 。你需要选择 恰好 一个下标（下标从 0 开始）并删除对应的元素。请注意剩下元素的下标可能会因为删除操作而发生改变。
 *
 * 比方说，如果 nums = [6,1,7,4,1] ，那么：
 *
 *
 * 选择删除下标 1 ，剩下的数组为 nums = [6,7,4,1] 。
 * 选择删除下标 2 ，剩下的数组为 nums = [6,1,4,1] 。
 * 选择删除下标 4 ，剩下的数组为 nums = [6,1,7,4] 。
 *
 *
 * 如果一个数组满足奇数下标元素的和与偶数下标元素的和相等，该数组就是一个 平衡数组 。
 *
 * 请你返回删除操作后，剩下的数组 nums 是 平衡数组 的 方案数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,1,6,4]
 * 输出：1
 * 解释：
 * 删除下标 0 ：[1,6,4] -> 偶数元素下标为：1 + 4 = 5 。奇数元素下标为：6 。不平衡。
 * 删除下标 1 ：[2,6,4] -> 偶数元素下标为：2 + 4 = 6 。奇数元素下标为：6 。平衡。
 * 删除下标 2 ：[2,1,4] -> 偶数元素下标为：2 + 4 = 6 。奇数元素下标为：1 。不平衡。
 * 删除下标 3 ：[2,1,6] -> 偶数元素下标为：2 + 6 = 8 。奇数元素下标为：1 。不平衡。
 * 只有一种让剩余数组成为平衡数组的方案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,1,1]
 * 输出：3
 * 解释：你可以删除任意元素，剩余数组都是平衡数组。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：0
 * 解释：不管删除哪个元素，剩下数组都不是平衡数组。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package jzoffer

// @lc code=start
// 1. The parity of the indices after the removed element changes.
// 2. Calculate prefix sums for even and odd indices
// separately to calculate for each index in O(1).
func waysToMakeFair1(nums []int) (res int) {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	even, odd, e, o := make([]int, n+1), make([]int, n+1), 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			o += v
		} else {
			e += v
		}
		even[i+1] = e
		odd[i+1] = o
	}

	// if even[n] == odd[n] {// 错误结论
	// 	return 0
	// }

	// 0  1  2  3
	// 奇 偶 奇 偶
	for i := range nums {
		// 删除后数组中全部奇数下标元素和为 even[i]+odd[n]-o
		// 全部偶数下标元素和为 odd[i]+even[n]-e
		// 若两者相等则说明删除下标 ii 后的数组为「平衡数组」。
		e, o := even[i+1], odd[i+1]
		if even[i]+odd[n]-o == odd[i]+even[n]-e { //删除偶数 奇数+偶 偶+奇
			res++
		}
	}
	return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/ways-to-make-a-fair-array/solution/sheng-cheng-ping-heng-shu-zu-de-fang-an-0mkaj/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func waysToMakeFair2(nums []int) (res int) {
	var preOdd, preEven, sufOdd, sufEven int
	for i, num := range nums {
		if i&1 > 0 {
			sufOdd += num
		} else {
			sufEven += num
		}
	}
	for i, num := range nums {
		if i&1 > 0 {
			sufOdd -= num
		} else {
			sufEven -= num
		}
		// num 之前的偶数和+之后的奇数和 == 之前的奇数和+之后的偶数和
		if preOdd+sufEven == preEven+sufOdd {
			res++
		}
		if i&1 > 0 {
			preOdd += num
		} else {
			preEven += num
		}
	}
	return
}

// 题解中的 odd1 even1 odd2 even2 这四个变量可以化简成一个变量。
// 把 odd1 + even2 == odd2 + even1 变形得 odd1 - even1 == odd2 - even2，
// 这样可以化简成两个变量 s1 = odd1 - even1 和 s2 = odd2 - even2，如果 s1 == s2 答案就加一。s1 和 s2 的更新就是 i 为奇数加，为偶数减。
// 再变形成 s1 - s2 == 0，那么令 s = s1 - s2，就可以化简成一个变量了，当 s == 0 时答案加一。
func waysToMakeFair(nums []int) (res int) {
	s := 0
	for i, v := range nums {
		if i&1 > 0 {
			s += v
		} else {
			s -= v
		}
	}
	for i, v := range nums {
		if i&1 > 0 {
			s -= v
		} else {
			s += v
		}
		if s == 0 {
			res++
		}
		if i&1 > 0 {
			s -= v
		} else {
			s += v
		}
	}
	return
}

// @lc code=end
