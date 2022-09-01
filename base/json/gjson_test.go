package base

import (
	"fmt"
	"testing"

	json "github.com/json-iterator/go"
)

// 默认仅对导出字段编码,
func TestJson1(t *testing.T) {
	v := person{Name: "1", Age: 2, Gender: "3", id: "4"}
	sv, err := json.Marshal(v)
	if err != nil {
		t.Error(err, "encoding err")
	}
	fmt.Println(string(sv))
	// {"Name":"1","Age":2,"Gender":"3"}

	s := `{"name":"tony","age":20,"gender":"male", "id":"xxx-xxx-xxx-xxx"}`
	var p1 person
	err = json.Unmarshal([]byte(s), &p1)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return
	}
	fmt.Printf("%#v\n", p1)
	// main.person{Name:"tony", Age:20, Gender:"male", id:""}
}

// 空切片和nil切片不同
func TestCodingSlice1(t *testing.T) {
	var nilSlice []int
	emptySlice := make([]int, 0, 5)

	println(nilSlice == nil)                              // true
	println(emptySlice == nil)                            // false
	println(nilSlice, len(nilSlice), cap(nilSlice))       // [0/0]0x0 0 0
	println(emptySlice, len(emptySlice), cap(emptySlice)) // [0/5]0xc00001e150 0 5

	snil, _ := json.Marshal(nilSlice)
	sEmpty, _ := json.Marshal(emptySlice)
	println(string(snil), string(sEmpty)) // null []

	m := map[string][]int{"nilSlice": nilSlice, "emptySlice": emptySlice}
	b, _ := json.Marshal(m)
	println(string(b)) // {"emptySlice":[],"nilSlice":null}
}

// 字节切片可能被编码为base64编码的文本
func TestCodingText1(t *testing.T) {
	m := map[string]interface{}{
		"byteSlice": []byte("hello, go"),
		"string":    "hello, go"}
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
	// {"byteSlice":"aGVsbG8sIGdv","string":"hello, go"}}
	/* 参考: go语言精进之路
	字节切片可以存储任意字节序列，可能会包含控制字符、字符“\0”及不合法Unicode字符等无法显示
	或导致乱码的内容。如果你的字节切片中存储的仅是合法Unicode字符的utf-8编码字节，
	又不想将其编码为base64输出，那么可以先将其转换为string类型后再用json包进行编码处理。
	*/
}
