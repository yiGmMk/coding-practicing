package base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	val := []interface{}{
		0,
		[]int{1},
		[]string{"go"},
		"gopher",
		struct{ A string }{A: "struct with string"},
	}

	var i interface{}
	for index, v := range val {
		i = v
		iVal := reflect.ValueOf(i)
		iType := reflect.TypeOf(i)
		fmt.Printf("%d.type:%+v\tval.type:%+v\tequal:%t\n",
			index, iType, iVal.Type(), reflect.DeepEqual(iType, iVal.Type()))
	}
}

func Add(i, j int) int {
	return i + j
}

type Calculator struct{}

func (c Calculator) Add(i, j int) int {
	return i + j
}

func TestCall(t *testing.T) {

	// 函数调用
	f := reflect.ValueOf(Add)
	var i = 5
	var j = 6
	vals := []reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)}
	ret := f.Call(vals)
	fmt.Println(ret[0].Int()) // 11

	// 方法调用
	c := reflect.ValueOf(Calculator{})
	m := c.MethodByName("Add")
	ret = m.Call(vals)
	fmt.Println(ret[0].Int()) // 11
}
