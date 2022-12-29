/*
 * @lc app=leetcode.cn id=2032 lang=golang
 *
 * [2032] 至少在两个数组中出现的值
 *
 * https://leetcode.cn/problems/two-out-of-three/description/
 *
 * algorithms
 * Easy (66.57%)
 * Likes:    51
 * Dislikes: 0
 * Total Accepted:    23.1K
 * Total Submissions: 32K
 * Testcase Example:  '[1,1,3,2]\n[2,3]\n[3]'
 *
 * 给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 元素各不相同的 数组，且由 至少 在 两个
 * 数组中出现的所有值组成。数组中的元素可以按 任意 顺序排列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
 * 输出：[3,2]
 * 解释：至少在两个数组中出现的所有值为：
 * - 3 ，在全部三个数组中都出现过。
 * - 2 ，在数组 nums1 和 nums2 中出现过。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
 * 输出：[2,3,1]
 * 解释：至少在两个数组中出现的所有值为：
 * - 2 ，在数组 nums2 和 nums3 中出现过。
 * - 3 ，在数组 nums1 和 nums2 中出现过。
 * - 1 ，在数组 nums1 和 nums3 中出现过。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
 * 输出：[]
 * 解释：不存在至少在两个数组中出现的值。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length, nums3.length <= 100
 * 1 <= nums1[i], nums2[j], nums3[k] <= 100
 *
 *
 */
package jzoffer

import (
	"math/bits"

	"github.com/samber/lo"
)

// @lc code=start

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1, m2, m3 := make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{})
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	res := make(map[int]struct{})
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}
	for _, v := range nums3 {
		m3[v] = struct{}{}
	}
	for _, v := range nums1 {
		if _, ok := m2[v]; ok {
			res[v] = struct{}{}
		}
		if _, ok := m3[v]; ok {
			res[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			res[v] = struct{}{}
		}
		if _, ok := m3[v]; ok {
			res[v] = struct{}{}
		}
	}
	for _, v := range nums3 {
		if _, ok := m1[v]; ok {
			res[v] = struct{}{}
		}
		if _, ok := m2[v]; ok {
			res[v] = struct{}{}
		}
	}
	vs := []int{}
	for k := range res {
		vs = append(vs, k)
	}
	return vs
}

func twoOutOfThree1(nums1 []int, nums2 []int, nums3 []int) []int {
	v12 := lo.Intersect(nums1, nums2)
	v23 := lo.Intersect(nums2, nums3)
	v13 := lo.Intersect(nums1, nums3)
	res := lo.Union(v12, v23, v13)
	return lo.FindUniques(res)
}

// 我们可以用「哈希表」来实现——由于只有三个数组，
// 所以我们一个整数的最低三个二进制位来标记某一个数在哪几个数组中，
// 1 表示该数在对应的数组中的，反之 0 表示不在。
// 最后我们只需要判断每一个数对应的标记数字中二进制位个数是否大于 1 即可
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/two-out-of-three/solution/zhi-shao-zai-liang-ge-shu-zu-zhong-chu-x-5131/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func twoOutOfThree2(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	mask := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mask[x] |= 1 << i
		}
	}
	for x, m := range mask {
		if m&(m-1) > 0 { // 二进制中
			ans = append(ans, x)
		}
	}
	return
}

func twoOutOfThree3(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	mask := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mask[x] |= 1 << i
		}
	}
	for x, m := range mask {
		// 二进制中
		if bits.OnesCount(uint(m)) > 1 {
			ans = append(ans, x)
		}
	}
	return
}

// @lc code=end
