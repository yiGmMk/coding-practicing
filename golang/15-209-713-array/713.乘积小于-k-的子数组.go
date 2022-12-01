/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于 K 的子数组
 *
 * https://leetcode.cn/problems/subarray-product-less-than-k/description/
 *
 * algorithms
 * Medium (48.82%)
 * Likes:    577
 * Dislikes: 0
 * Total Accepted:    82K
 * Total Submissions: 168K
 * Testcase Example:  '[10,5,2,6]\n100'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,5,2,6], k = 100
 * 输出：8
 * 解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
 * 需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 0
 * 输出：0
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * 1 <= nums[i] <= 1000
 * 0 <= k <= 10^6
 *
 *
 */
package jzoffer

import (
	"math"
	"sort"
)

// @lc code=start
// TODO: 滑动窗口
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/ZVAVXX/solution/cheng-ji-xiao-yu-k-de-zi-shu-zu-by-leetc-xqx8/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}

// 暴力求解
func numSubarrayProductLessThanK2(nums []int, k int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for r, product := i, nums[i]; r < n && product < k; {
			ans++
			r++
			if r == n {
				break
			}
			product *= nums[r]
		}
	}
	return
}

/*
作者：cranky-ptolemylu3
链接：https://leetcode.cn/problems/ZVAVXX/solution/by-cranky-ptolemylu3-1wjl/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

public int numSubarrayProductLessThanK(int[] nums, int k) {
    //结果
	int res = 0;
	int len = nums.length;
    //i作为左窗口
	for (int i = 0; i < len; i++) {
		int r = i;
		int sum = nums[i];
		while (sum < k && r < len) {
            //符合就+1计数
			res++;
            //右窗口往右滑动
			r++;
            //防止数据索引超标
			if(r >= len ) break;
			sum = sum * nums[r];
		}
	}
	return res;
}
*/

// 二分查找,直接用元素乘积可能溢出,两边取对数,可以避免溢出
// 乘法通过取对数变成加法
func numSubarrayProductLessThanKBinary(nums []int, k int) (ans int) {
	if k == 0 {
		return
	}
	n := len(nums)
	logPrefix := make([]float64, n+1)
	for i, num := range nums {
		logPrefix[i+1] = logPrefix[i] + math.Log(float64(num))
	}
	logK := math.Log(float64(k))
	for j := 1; j <= n; j++ {
		l := sort.SearchFloat64s(logPrefix[:j], logPrefix[j]-logK+1e-10)
		ans += j - l
	}
	return
}

// @lc code=end
