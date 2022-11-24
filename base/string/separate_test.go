package base

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

var (
	_cutTestCase = []struct {
		in            string
		sept          string
		before, after string
		found         bool
	}{
		{
			in:     "yigmmk@github.com",
			sept:   "@",
			before: "yigmmk",
			after:  "github.com",
			found:  true,
		},
		{
			found: true,
		},
	}

	_fieldTestCase = []struct {
		in   string
		out  []string
		sept rune
	}{
		{
			in:   "china,japan,korean,US,UK",
			out:  []string{"china", "japan", "korean", "US", "UK"},
			sept: ',',
		},
		{
			in:   "china japan\nkorean\tUS UK",
			out:  []string{"china", "japan", "korean", "US", "UK"},
			sept: ' ',
		},
	}

	_spiltTestCase = []struct {
		in   string
		out  []string
		sept string
	}{
		{
			in:   "china,japan,korean,US,UK",
			out:  []string{"china", "japan", "korean", "US", "UK"},
			sept: ",",
		},
		{
			in:   "Go社区欢迎你",
			out:  []string{"G", "o", "社", "区", "欢", "迎", "你"},
			sept: "",
		},
	}
)

func TestSeparate(t *testing.T) {
	t.Run("cut", func(t *testing.T) {
		for _, v := range _cutTestCase {
			before, after, found := strings.Cut(v.in, v.sept)
			if !(before == v.before && after == v.after && v.found == found) {
				t.Errorf("cut %s error\nbefore:%s\nafter:%s\nfound:%+v",
					v.in, v.before, v.after, v.found)
			}
		}
	})

	// Fields函数采用了Unicode空白字符的定义，下面的字符均会被识别为空白字符
	// '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
	t.Run("field", func(t *testing.T) {
		for i, v := range _fieldTestCase {
			var out []string
			if unicode.IsSpace(v.sept) {
				out = strings.Fields(v.in)
			} else {
				out = strings.FieldsFunc(v.in, func(r rune) bool { return r == rune(v.sept) })
			}
			if !reflect.DeepEqual(v.out, out) && !(v.out == nil || out == nil) {
				t.Errorf("field[%d] %s error", i, v.in)
			}
		}
	})

}

func TestSpilt(t *testing.T) {
	t.Run("spilt", func(t *testing.T) {
		for i, v := range _spiltTestCase {
			out := strings.Split(v.in, string(v.sept))
			if !reflect.DeepEqual(v.out, out) && !(v.out == nil || out == nil) {
				t.Errorf("field[%d] %s error", i, v.in)
			}
		}
	})
}

// 使用Split相关函数分割字符串
// fmt.Printf("%q\n", strings.Split("a,b,c", ","))    // ["a" "b" "c"]
// fmt.Printf("%q\n", strings.Split("a,b,c", "b"))    // ["a," ",c"]
// fmt.Printf("%q\n", strings.Split("Go社区欢迎你", ""))    // ["G" "o" "社" "区" "欢" "迎" "你"]
// fmt.Printf("%q\n", strings.Split("abc", "de"))    // ["abc"]
// fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 2))    // ["a" "b,c,d"]
// fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 3))    // ["a" "b" "c,d"]
// fmt.Printf("%q\n", strings.SplitAfter("a,b,c,d", ","))    // ["a," "b," "c," "d"]
// fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d", ",", 2))    // ["a," "b,c,d"]

// 使用Split相关函数分割字节切片
//fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))    // ["a" "b" "c"]
//fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte("b")))    // ["a," ",c"]
//fmt.Printf("%q\n", bytes.Split([]byte("Go社区欢迎你"), nil))    // ["G" "o" "社" "区" "欢" "迎" "你"]
//fmt.Printf("%q\n", bytes.Split([]byte("abc"), []byte("de")))    // ["abc"]
//fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c,d"), []byte(","), 2))    // ["a" "b,c,d"]
//fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c,d"), []byte(","), 3))    // ["a" "b" "c,d"]
//fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c,d"), []byte(",")))    // ["a," "b," "c," "d"]
//fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c,d"), []byte(","), 2))    // ["a," "b,c,d"]
