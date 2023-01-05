/*
 * @lc app=leetcode.cn id=1802 lang=golang
 *
 * [1802] 有界数组中指定下标处的最大值
 *
 * https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/description/
 *
 * algorithms
 * Medium (29.20%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    10.6K
 * Total Submissions: 31.5K
 * Testcase Example:  '4\n2\n6'
 *
 * 给你三个正整数 n、index 和 maxSum 。你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：
 *
 *
 * nums.length == n
 * nums[i] 是 正整数 ，其中 0 <= i < n
 * abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
 * nums 中所有元素之和不超过 maxSum
 * nums[index] 的值被 最大化
 *
 *
 * 返回你所构造的数组中的 nums[index] 。
 *
 * 注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 4, index = 2,  maxSum = 6
 * 输出：2
 * 解释：数组 [1,1,2,1] 和 [1,2,2,1] 满足所有条件。不存在其他在指定下标处具有更大值的有效数组。
 *
 *
 * 示例 2：
 *
 * 输入：n = 6, index = 1,  maxSum = 10
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= maxSum <= 10^9
 * 0 <= index < n
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
// 提示
// 1.如何生成一个有效的数组使nums[index]==target?
// 2.使nums[index]=target,nums[index-i]=target-i,nums[index+i]=target-i =>检查sum<=maxSum是否成立?
// 3.使用1+2+...+n=n*(n+1)/2 快速计算sum

// https://zxi.mytechroad.com/blog/algorithms/binary-search/leetcode-1802-maximum-value-at-a-given-index-in-a-bounded-array/
func maxValue0(n int, index int, maxSum int) int {
	l, r := 1, maxSum+1
	for l < r {
		m := l + (r-l)/2
		t1, t2 := m-index, m-(n-1-index) // m是最大值,求左右两边等差数列的最小值
		s1 := (t1 + m) * (index + 1) / 2
		s2 := (m + t2) * (n - index) / 2
		if t1 <= 0 {
			s1 = (1+m)*m/2 + (index - m + 1)
		}
		if t2 <= 0 {
			s2 = (1+m)*m/2 + (n - index - m)
		}
		if (s1 + s2 - m) > maxSum {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}

// https://zxi.mytechroad.com/blog/algorithms/binary-search/leetcode-1802-maximum-value-at-a-given-index-in-a-bounded-array/
func maxValue00(n int, index int, maxSum int) int {
	l, r := 0, maxSum+1
	for l+1 < r {
		m := l + (r-l)/2
		t1, t2 := m-index, m-(n-1-index) // m是最大值,求左右两边等差数列的最小值
		s1 := (t1 + m) * (index + 1) / 2 // 等差数列求和
		s2 := (m + t2) * (n - index) / 2
		if t1 <= 0 {
			s1 = (1+m)*m/2 + (index - m + 1) // m<长度,一侧的值降到1后剩余的都是1,m项使用等差数列求和公式求和,剩余的都是1
		}
		if t2 <= 0 {
			s2 = (1+m)*m/2 + (n - index - m)
		}
		if (s1 + s2 - m) > maxSum {
			r = m
		} else {
			l = m
		}
	}
	return l
}

func maxValue(n int, index int, maxSum int) int {
	// 最大值可能是maxSum,这里查找范围+1
	return sort.Search(maxSum+1, func(m int) bool {
		t1, t2 := m-index, m-(n-1-index) // m是最大值,求左右两边等差数列的最小值
		s1 := (t1 + m) * (index + 1) / 2 // 等差数列求和
		s2 := (m + t2) * (n - index) / 2
		if t1 <= 0 {
			s1 = (1+m)*m/2 + (index - m + 1) // m<长度,一侧的值降到1后剩余的都是1,m项使用等差数列求和公式求和,剩余的都是1
		}
		if t2 <= 0 {
			s2 = (1+m)*m/2 + (n - index - m)
		}
		sum := s1 + s2 - m
		return sum > maxSum // s1,s2里多加了个m
	}) - 1 // 查找返回的是第一个>maxSum的,减一的是<=maxSum的最大值
}

// 假设nums[idx] = m，则整体三角形面积为m平方，减去两角的小三角形就是总面积
// s1的面积为1 + 2 + ... + m-(idx+1)，等差数列
// s2同理，1 + 2 + ... + m-(n-idx)
// 二分高度m即可
// https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/solution/you-jie-shu-zu-zhong-zhi-ding-xia-biao-c-aav4/
func maxValue1(n int, index int, maxSum int) int {
	l, r := -1, maxSum+1
	maxSum -= n
	for l+1 < r {
		mid := (l + r) / 2
		x, y := mid-index-1, mid-n+index
		s, s1, s2 := mid*mid, max(0, x)*(x+1)/2, max(0, y)*(y+1)/2
		if s-s1-s2 <= maxSum {
			l = mid
		} else {
			r = mid
		}
	}
	return l + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 计算单边元素和
// 等差数列求和 (a1+an)*n/2 ,即(首项+末项)*项数/2
func f(big, length int) int {
	if length == 0 {
		return 0
	}
	if length <= big {
		return (2*big + 1 - length) * length / 2 // 未降到1,
	}
	return (big+1)*big/2 + length - big // length>big时,已经降到1,等差数列求和
}

func maxValue2(n, index, maxSum int) int {
	return sort.Search(maxSum, func(mid int) bool {
		left := index          // 左边长度,不包含index
		right := n - index - 1 // 右边长度,不包含index
		// 和=nums[index]+nums[index]左边部分之和+nums[index]右边部分之和
		sl, sr := f(mid, left), f(mid, right)
		return mid+sl+sr >= maxSum
	})
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/solution/you-jie-shu-zu-zhong-zhi-ding-xia-biao-c-aav4/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end
