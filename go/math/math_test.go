package math

import "testing"

// 2的次幂
func Test1(t *testing.T) {
	check := func(x int) bool {
		if x > 0 && x&(x-1) == 0 {
			return true
		}
		return false
	}

	i, n := 1, 30
	for n > 0 {
		if !check(i) {
			t.Errorf("check 错误,i=%d,n=%d", i, n)
		}
		i *= 2
		n--
	}
}
