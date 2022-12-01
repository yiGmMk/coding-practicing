package base

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

// string的for range每次返回一个rune,即utf-8的一个'码点',不是一个字节.
// 按下标访问才是字节.

// go的string是utf-8编码(变长编码),占用1~4字节,参考:<go语言精进之路>https://weread.qq.com/web/reader/b8f32d2072895edbb8fbb04k764323602597647966b7a1c
// string是unicode的字符串
// 按下标访问得到的是byte,字节码,如果是中文,下标访问拿不到完整的字符数据(中文3字节,下标访问得到的是byte,1字节)
// for range 访问得到的是rune,4字节
func TestStrRange(t *testing.T) {

	for _, s := range "Hi,中国" {
		fmt.Printf("0x%X\n", s)
	}

	str := `中\a建`
	fmt.Println(unsafe.Sizeof(rune('中')), unsafe.Sizeof(str[0]), str[0], string(str[0]))

	// 按占用空间
	row := 0
	fmt.Printf("-----------------%d-----------------------\n", row)
	for i := 0; i < len(str); i++ {
		v := str[i]
		fmt.Println(i, ":", len(str), v, string(v), rune(v))
	}

	// range
	row++
	fmt.Printf("-----------------%d-----------------------\n", row)
	for i, v := range str {
		fmt.Println(i, ":", len(str), v, string(v), rune(v))
	}

	// range 底层数组
	row++
	fmt.Printf("-----------------%d-----------------------\n", row)
	for i, v := range []byte(str) {
		fmt.Println(i, ":", len(str), v, string(v), rune(v))
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

// 参考: https://golang.google.cn/blog/strings
// Go source code is always UTF-8.
// A string holds arbitrary bytes. 任意字节
// A string literal, absent(缺席,不在场的) byte-level escapes(逃走,转义), always holds valid UTF-8 sequences.
// Those sequences represent Unicode code points, called runes.(Unicode用code point表示单个字符)
// No guarantee is made in Go that characters in strings are normalized.
func TestString(t *testing.T) {
	var i int
	index := func() string {
		i++
		return strconv.Itoa(i) + "."
	}
	fString := func(str string) {
		fmt.Println(index(), str) // 1.

		fmt.Printf("%s", index()) // 2.
		for i := 0; i < len(str); i++ {
			fmt.Printf("%x ", str[i])
		}
		fmt.Println()

		fmt.Printf("%s%x\n", index(), str)  // 按十六进制打印
		fmt.Printf("%s% x\n", index(), str) // 十六进制加空格打印

		// There’s more. The %q (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.
		// %q (引号)谓词将转义字符串中任何不可打印的字节序列，因此输出是明确的。
		// %+q 这个标志使得输出不仅转义不可打印的序列，而且转义任何非 ASCII 字节，所有这些都是在解释 UTF-8时进行的。
		fmt.Printf("%s%q\n", index(), str)
		fmt.Printf("%s%+q\n", index(), str)
	}

	fByteSlice := func(str []byte) {
		fmt.Println(index(), str) // 1.

		fmt.Printf("%s", index()) // 2.
		for i := 0; i < len(str); i++ {
			fmt.Printf("%x ", str[i])
		}
		fmt.Println()

		fmt.Printf("%s%x\n", index(), str)  // 按十六进制打印
		fmt.Printf("%s% x\n", index(), str) // 十六进制加空格打印

		// There’s more. The %q (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.
		// %q (引号)谓词将转义字符串中任何不可打印的字节序列，因此输出是明确的。
		// %+q 这个标志使得输出不仅转义不可打印的序列，而且转义任何非 ASCII 字节，所有这些都是在解释 UTF-8时进行的。
		fmt.Printf("%s%q\n", index(), str)
		fmt.Printf("%s%+q\n", index(), str)
	}

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" // ⌘ => UTF-8的十六进制,用byte表示为e2 8c 98
	fString(sample)

	fmt.Println("-------byte---------")
	fByteSlice([]byte(sample))

	fmt.Println("---UTF-8 and string literals----")
	// As we saw, indexing a string yields its bytes, not its characters:
	// a string is just a bunch of bytes. That means that when we store a character value in a string,
	// we store its byte-at-a-time representation.
	// 按下标访问得到字符串底层存储的byte,不是字符
	// len(string)是占用空间/字节
	// Source code in Go is defined to be UTF-8 text; no other representation is allowed
	// go源代码是UTF-8的,且只支持UTF-8表示
	const placeOfInterest = `⌘`
	fmt.Printf("%slen:%d,plain string: ", index(), len(placeOfInterest))
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("%squoted string: ", index())
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("%shex bytes: ", index())
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	// 使用 range 遍历
	fmt.Printf("%sfor range\n", index())
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	// 标准库中的辅助函数 utf
	fmt.Printf("%sutf package\n", index())
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func TestIo(t *testing.T) {

	var buf bytes.Buffer
	var s = "I love Go!!"

	_, err := io.Copy(&buf, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String()) // "I love Go!!"

	buf.Reset()
	var b = []byte("I love Go!!")
	_, err = io.Copy(&buf, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String()) // "I love Go!!"

}

func TestLen(t *testing.T) {
	s := "大家好"
	fmt.Printf("字符串\"%s\"的长度为%d,字符个数:%d\n",
		s, len(s), utf8.RuneCountInString(s)) // 长度为9
}
