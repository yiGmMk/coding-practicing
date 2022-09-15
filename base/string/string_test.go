package base

import (
	"log"
	"testing"
	"unsafe"
)

// string是unicode的字符串
// 按下标访问得到的是byte,字节码,如果是中文,下标访问拿不到完整的字符数据(中文4字节,下标访问得到的是byte,1字节)
// for range 访问得到的是rune,4字节
func TestStrRange(t *testing.T) {
	str := `中\abcd,./会等凌绝顶,一览众山小\`
	log.Println(unsafe.Sizeof(rune('中')), unsafe.Sizeof(str[0]), str[0], string(str[0]))

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
