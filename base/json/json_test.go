package base

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// chapter10/sources/go-trap/json_1.go
type person struct {
	Name   string
	Age    int
	Gender string
	id     string
}

// 默认仅对导出字段编码,
func TestDefaultJson(t *testing.T) {
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
func TestCodingSlice(t *testing.T) {
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
func TestCodingText(t *testing.T) {
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

// interface{}的底层类型是float64，而不是int。
// 当JSON文本中的整型数值被解码为interface{}类型时，其底层真实类型为float64
// json包提供了Number类型来存储JSON文本中的各类数值类型，并可以转换为整型（int64）、浮点型（float64）及字符串
func TestJosnInterface(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic:", e)
		}
	}()
	flexibleJsonUnmarshal := func() {
		s := `{"age": 23, "name": "tony"}`

		m := map[string]interface{}{}
		_ = json.Unmarshal([]byte(s), &m)

		age := m["age"].(int)
		// panic: interface conversion: interface {} is float64, not int
		fmt.Println("age =", age)
	}
	flexibleJsonUnmarshal()
}

// json包提供了Number类型来存储JSON文本中的各类数值类型，并可以转换为整型（int64）、浮点型（float64）及字符串
func TestFlexibleJsonUnmarshalImproved(t *testing.T) {
	flexibleJsonUnmarshalImproved := func() {
		s := `{"age": 23, "name": "tony", "cash": 33.33}`
		m := map[string]interface{}{}
		d := json.NewDecoder(strings.NewReader(s))
		d.UseNumber()

		_ = d.Decode(&m)
		age, _ := m["age"].(json.Number).Int64()
		fmt.Println("age =", age)
		// age = 23

		cash, _ := m["cash"].(json.Number).Float64()
		fmt.Println("cash =", cash)
		// cash = 33.33
	}
	flexibleJsonUnmarshalImproved()
}
