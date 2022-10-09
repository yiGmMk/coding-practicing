package singleton

import "testing"

func TestSingleton(t *testing.T) {
	cnt := 10
	var uid string
	for i := 0; i < cnt; i++ {
		i := EagerSingleton()
		if uid == "" {
			uid = i.Uid
		}
		if i.Uid != uid {
			t.Error("not equal")
		}
	}
}

func TestLzaySingleton(t *testing.T) {
	cnt := 10
	var uid string
	for i := 0; i < cnt; i++ {
		i := LazySingleton()
		if uid == "" {
			uid = i.Uid
		}
		if i.Uid != uid {
			t.Error("not equal")
		}
	}
}

/*
 *goos: linux
 *goarch: amd64
 *pkg: github.com/yiGmMk/leetcode/design-pattern/creational/singleton
 *cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
 *BenchmarkLazy-2   	466875750	         2.546 ns/op	       0 B/op	       0 allocs/op
 *PASS
 **/
func BenchmarkLazy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LazySingleton()
	}
}

/*
 *goos: linux
 *goarch: amd64
 *pkg: github.com/yiGmMk/leetcode/design-pattern/creational/singleton
 *cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
 *BenchmarkEager-2   	1000000000	         0.3653 ns/op	       0 B/op	       0 allocs/op
 *PASS
 */
func BenchmarkEager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EagerSingleton()
	}
}
