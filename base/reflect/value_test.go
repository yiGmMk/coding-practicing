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

func TestAlter(t *testing.T) {

	type Person struct {
		Name string
		age  int
	}

	var val reflect.Value
	var n = 17
	t.Run("int", func(t *testing.T) {

		fmt.Println("int:")
		val = reflect.ValueOf(n)
		fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
			val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	})

	t.Run("*int", func(t *testing.T) {
		fmt.Println("\n*int:")
		val = reflect.ValueOf(&n)
		fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
			val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
		val = reflect.ValueOf(&n).Elem()
		fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
			val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true
	})

	t.Run("slice", func(t *testing.T) {
		fmt.Println("\nslice:")
		var sl = []int{5, 6, 7}
		val = reflect.ValueOf(sl)
		fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
			val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
		val = val.Index(0)
		fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
			val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true
	})

	fmt.Println("\narray:")
	var arr = [3]int{5, 6, 7}
	val = reflect.ValueOf(arr)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Index(0)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true

	fmt.Println("\nptr to array:")
	var pArr = &[3]int{5, 6, 7}
	val = reflect.ValueOf(pArr)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true
	val = val.Index(0)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true

	fmt.Println("\nstruct:")
	p := Person{"tony", 33}
	val = reflect.ValueOf(p)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val1 := val.Field(0) // Name
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val1.CanSet(), val1.CanAddr(), val1.CanInterface()) // false false true
	val2 := val.Field(1) // age
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val2.CanSet(), val2.CanAddr(), val2.CanInterface()) // false false false

	fmt.Println("\nptr to struct:")
	pp := &Person{"tony", 33}
	val = reflect.ValueOf(pp)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true
	val1 = val.Field(0) // Name
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val1.CanSet(), val1.CanAddr(), val1.CanInterface()) // true true true
	val2 = val.Field(1) // age
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val2.CanSet(), val2.CanAddr(), val2.CanInterface()) // false true false

	fmt.Println("\ninterface:")
	var i interface{} = &Person{"tony", 33}
	val = reflect.ValueOf(i)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true

	fmt.Println("\nmap:")
	var m = map[string]int{
		"tony": 23,
		"jim":  34,
	}
	val = reflect.ValueOf(m)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true

	val.SetMapIndex(reflect.ValueOf("tony"), reflect.ValueOf(12))
	fmt.Println(m) // map[jim:34 tony:12]
}
