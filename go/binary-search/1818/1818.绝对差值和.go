/*
 * @lc app=leetcode.cn id=1818 lang=golang
 *
 * [1818] 绝对差值和
 *
 * https://leetcode.cn/problems/minimum-absolute-sum-difference/description/
 *
 * algorithms
 * Medium (37.51%)
 * Likes:    145
 * Dislikes: 0
 * Total Accepted:    30.4K
 * Total Submissions: 81K
 * Testcase Example:  '[1,7,5]\n[2,3,5]'
 *
 * 给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。
 *
 * 数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 ）的 总和（下标从 0 开始）。
 *
 * 你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。
 *
 * 在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。因为答案可能很大，所以需要对 10^9 + 7 取余 后返回。
 *
 * |x| 定义为：
 *
 *
 * 如果 x >= 0 ，值为 x ，或者
 * 如果 x  ，值为 -x
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,7,5], nums2 = [2,3,5]
 * 输出：3
 * 解释：有两种可能的最优方案：
 * - 将第二个元素替换为第一个元素：[1,7,5] => [1,1,5] ，或者
 * - 将第二个元素替换为第三个元素：[1,7,5] => [1,5,5]
 * 两种方案的绝对差值和都是 |1-2| + (|1-3| 或者 |5-3|) + |5-5| = 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
 * 输出：0
 * 解释：nums1 和 nums2 相等，所以不用替换元素。绝对差值和为 0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
 * 输出：20
 * 解释：将第一个元素替换为第二个元素：[1,10,4,4,2,7] => [10,10,4,4,2,7]
 * 绝对差值和为 |10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums1.length
 * n == nums2.length
 * 1
 * 1
 *
 *
 */
package jzoffer

import "sort"

// TODO,没思路
// @lc code=start
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

func minAbsoluteSumDiff(nums1 []int, nums2 []int) (res int) {
	rec := make([]int, len(nums1), len(nums1))
	copy(rec, nums1)
	sort.Ints(rec)
	var (
		n         = len(nums1)
		sum, maxn = 0, 0
		mod       = int(1e9 + 7)
	)
	for i := 0; i < n; i++ {
		diff := abs(nums1[i] - nums2[i])
		sum += diff

		j := sort.Search(n, func(j int) bool { return rec[j] >= nums2[i] })
		if j < n {
			maxn = max(maxn, diff-(rec[j]-nums2[i]))
		}
		if j > 0 {
			maxn = max(maxn, diff-(nums2[i]-rec[j-1]))
		}
	}
	return (sum - maxn) % mod
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/minimum-absolute-sum-difference/solution/jue-dui-chai-zhi-he-by-leetcode-solution-gv78/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// 题目要求使 |nums1[i] - nums2[i]|尽可能小,可以替换一个nums1中的
// 替换后对结果的改变为|nums1[i]-nums2[i]|-|nums1[j]-nums2[i]|,为了使结果小,则应该最大化这个差值
func minAbsoluteSumDiff1(nums1, nums2 []int) int {
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		j := rec.Search(v) // >=v的数
		if j < n {
			maxn = max(maxn, diff-(rec[j]-v))
		}
		if j > 0 { // <v的最大值
			maxn = max(maxn, diff-(v-rec[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

// @lc code=end
