package jzoffer

import "math"

// TODO:
func divideBinary(a, b int) int {
	if a == math.MinInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 { // 考虑除数为最小值的情况
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	if a == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if a > 0 {
		a = -a
		rev = !rev
	}
	if b > 0 {
		b = -b
		rev = !rev
	}

	candidates := []int{b}
	for y := b; y >= a-y; { // 注意溢出
		y += y
		candidates = append(candidates, y)
	}

	ans := 0
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] >= a {
			ans |= 1 << i
			a -= candidates[i]
		}
	}
	if rev {
		return -ans
	}
	return ans
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/xoh6Oh/solution/zheng-shu-chu-fa-by-leetcode-solution-3572/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
