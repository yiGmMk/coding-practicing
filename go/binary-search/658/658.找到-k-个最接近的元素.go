/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 *
 * https://leetcode.cn/problems/find-k-closest-elements/description/
 *
 * algorithms
 * Medium (48.11%)
 * Likes:    463
 * Dislikes: 0
 * Total Accepted:    83.1K
 * Total Submissions: 172.7K
 * Testcase Example:  '[1,2,3,4,5]\n4\n3'
 *
 * 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
 *
 * 整数 a 比整数 b 更接近 x 需要满足：
 *
 *
 * |a - x| < |b - x| 或者
 * |a - x| == |b - x| 且 a < b
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [1,2,3,4,5], k = 4, x = 3
 * 输出：[1,2,3,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,2,3,4,5], k = 4, x = -1
 * 输出：[1,2,3,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= arr.length
 * 1 <= arr.length <= 10^4
 * arr 按 升序 排列
 * -10^4 <= arr[i], x <= 10^4
 *
 *
 */
package jzoffer

import (
	"sort"
)

// @lc code=start
// TODO
// 首先将数组 arr 按照「更接近」的定义进行排序，如果 a 比 b 更接近 x，
// 那么 a 将排在 b 前面。排序完成之后，k 个最接近的元素就是数组
// arr 的前 k 个元素，将这 k 个元素从小到大进行排序后，直接返回。
func findClosestElements(arr []int, k int, x int) (res []int) {
	sort.SliceStable(arr, func(i, j int) bool {
		return abs(arr[i]-x) < abs(arr[j]-x)
	})
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/find-k-closest-elements/solution/zhao-dao-k-ge-zui-jie-jin-de-yuan-su-by-ekwtd/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func findClosestElements1(arr []int, k int, x int) (res []int) {
	// 使用二分查找将数组分割为两部分
	// L 0:i的值 <x
	// R i:n    >=x
	r := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
	l := r - 1
	for ; k > 0; k-- {
		if l < 0 {
			r++
		} else if r >= len(arr) || x-arr[l] <= arr[r]-x {
			l--
		} else {
			r++
		}
	}
	return arr[l+1 : r]
}

// @lc code=end
