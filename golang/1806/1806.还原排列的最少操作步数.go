/*
 * @lc app=leetcode.cn id=1806 lang=golang
 *
 * [1806] 还原排列的最少操作步数
 *
 * https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/description/
 *
 * algorithms
 * Medium (65.22%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    9.9K
 * Total Submissions: 13.9K
 * Testcase Example:  '2'
 *
 * 给你一个偶数 n​​​​​​ ，已知存在一个长度为 n 的排列 perm ，其中 perm[i] == i​（下标 从 0 开始 计数）。
 *
 * 一步操作中，你将创建一个新数组 arr ，对于每个 i ：
 *
 *
 * 如果 i % 2 == 0 ，那么 arr[i] = perm[i / 2]
 * 如果 i % 2 == 1 ，那么 arr[i] = perm[n / 2 + (i - 1) / 2]
 *
 *
 * 然后将 arr​​ 赋值​​给 perm 。
 *
 * 要想使 perm 回到排列初始值，至少需要执行多少步操作？返回最小的 非零 操作步数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：1
 * 解释：最初，perm = [0,1]
 * 第 1 步操作后，perm = [0,1]
 * 所以，仅需执行 1 步操作
 *
 * 示例 2：
 *
 *
 * 输入：n = 4
 * 输出：2
 * 解释：最初，perm = [0,1,2,3]
 * 第 1 步操作后，perm = [0,2,1,3]
 * 第 2 步操作后，perm = [0,1,2,3]
 * 所以，仅需执行 2 步操作
 *
 * 示例 3：
 *
 *
 * 输入：n = 6
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2
 * n​​​​​​ 是一个偶数
 *
 *
 */

package jzoffer

import "reflect"

// @lc code=start
// 对单个数进行模拟,0,n-1这两个位置不变的
// 通过观察可以发现每一步操作相当于
// 将前一半的元素扩散到偶数下标的位置，
// 后一半扩散到奇数下标的位置。问题其实就是perm中每个元素通过这样变化返回
// 原来状态的周期是多长，因此可以只对单个元素进行模拟即可。
// https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/solution/huan-yuan-pai-lie-de-zui-shao-cao-zuo-bu-d9cn/
func reinitializePermutation(n int) int {
	if n == 2 {
		return 1
	}
	step, idx := 0, 1
	for step == 0 || idx != 1 {
		if idx < n/2 {
			idx = 2 * idx
		} else {
			idx = 2*(idx-n/2) + 1 // 2*idx-(n-1)
		}
		step++
	}
	return step
}

func reinitializePermutation1(n int) int {
	perm, ori := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
		ori[i] = i
	}

	arr, res := make([]int, n), 0

	for {
		for i := 0; i < n; i++ {
			// i%2==0 arr[i]=perm[i/2]
			// i%2==1 arr[i]=perm[n/2+(i-1)/2]
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else if i%2 == 1 {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		res++
		copy(perm, arr)
		if reflect.DeepEqual(ori, perm) {
			return res
		}
	}
}

// @lc code=end
