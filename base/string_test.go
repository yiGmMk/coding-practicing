package base

import (
	"testing"
)

func TestStrRange(t *testing.T) {
	str := `\abcd,./会等凌绝顶,一览众山小\`
	for i := 0; i < len(str); i++ {
		v := str[i]
		println(i, ":", len(str), v, string(v), rune(v))
	}

	//
	for i, v := range str {
		println(i, ":", len(str), v, string(v), rune(v))
	}
}

// string 底层结构里存了长度,len耗时为O(1)
func BenchmarkLenStr(b *testing.B) {
	str := `\abcd,./会等凌绝顶,一览众山小\`
	for i := 0; i < b.N; i++ {
		_ = len(str)
	}
}

func BenchmarkLenStr1(b *testing.B) {
	str := `\abcd,./会等凌绝顶,一览众山小\`
	l := len(str)
	for i := 0; i < b.N; i++ {
		_ = l
	}
}
