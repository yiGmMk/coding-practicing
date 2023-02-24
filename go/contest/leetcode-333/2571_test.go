package jzoffer

import (
	"fmt"
	"testing"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/* 判断一个数是否是2的幂次: n > 0 && n & (n-1) == 0 => (没有在同一位置都为1的)
 * 2的幂次的二进制中1的个数是1=> (0)1 (2)10 (4)100 (8)1000
 */
func minOperations1(n int) (res int) {
	if n&(n-1) == 0 {
		return 1
	}
	for n > 0 {
		i := 1
		for ; i*2 <= n; i *= 2 {
		}
		next := i * 2
		if abs(n-next) < abs(n-i) {
			n = abs(n - next)
		} else {
			n = abs(n - i)
		}

		res++
	}
	return
}

/*
低位的1会影响高位的1,优先消除低位的1再消除高位的,当有连续的1时优先加分(可以消除多个1)

作者：endlesscheng
链接：https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0/solution/ji-yi-hua-sou-suo-by-endlesscheng-cm6l/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func minOperations(n int) int {
	ans := 1
	for n&(n-1) > 0 { // n 不是 2 的幂次
		lb := n & -n // lowbit
		fmt.Printf("n=%b,-n=%b,lb=%b\n", n, -n, lb)
		if n&(lb<<1) > 0 { // 多个连续 1
			n += lb
		} else {
			n -= lb // 单个 1
		}
		ans++
	}
	return ans
}

var _testcase = []struct {
	n, expect int
}{
	{
		n:      39,
		expect: 3,
	},
	{
		expect: 3,
		n:      54,
	},
}

// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0/solution/ji-yi-hua-sou-suo-by-endlesscheng-cm6l/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
var cache = map[int]int{}

func minOperations2(n int) int {
	if n&(n-1) == 0 { // n 是 2 的幂次
		return 1
	}
	if res, ok := cache[n]; ok {
		return res
	}
	lb := n & -n
	res := 1 + min(minOperations2(n+lb), minOperations2(n-lb))
	cache[n] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test2571(t *testing.T) {
	for i, v := range _testcase {
		got := minOperations1(v.n)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}

		got = minOperations(v.n)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}

		got = minOperations2(v.n)
		if got != v.expect {
			t.Errorf("[%d].got:%d,expect:%d", i, got, v.expect)
		}
	}
}
