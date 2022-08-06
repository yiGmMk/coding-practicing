package jzoffer

// n*log(n)
func countBitsOLogn(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = countOne(i)
	}
	return result
}

// log(n)
func countOne(num int) int {
	res := 0
	for num > 0 {
		res += num & 1
		num = num >> 1
	}
	return res
}

// O(n)
func countBitsOn(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}

// O(n)
// 动态规划+最低有效位
func countBitsLowBit(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

// O(n*log(n))
// 利用Brian Kernighan算法，可以在一定程度上进一步提升计算速度。
func countBitsBrianKernighan(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/counting-bits/solution/bi-te-wei-ji-shu-by-leetcode-solution-0t1i/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
