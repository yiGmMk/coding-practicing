package base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	var pi = (*int)(nil)
	var ps = (*string)(nil)
	typ := reflect.TypeOf(pi)
	fmt.Println(typ.Kind(), typ.String()) // ptr *int

	typ = reflect.TypeOf(ps)
	fmt.Println(typ.Kind(), typ.String()) // ptr *string
}

func TestType2(t *testing.T) {
	var (
		val reflect.Value
		typ reflect.Type
	)
	// 原生复合类型
	t.Run("slice", func(t *testing.T) {
		var sl = []int{5, 6} // 切片
		val = reflect.ValueOf(sl)
		typ = reflect.TypeOf(sl)
		fmt.Printf("[%d %d]\n", val.Index(0).Int(),
			val.Index(1).Int()) // [5, 6]
		fmt.Println(typ.Kind(), typ.String()) // slice []int
	})

	t.Run("array", func(t *testing.T) {
		var arr = [3]int{5, 6} // 数组
		val = reflect.ValueOf(arr)
		typ = reflect.TypeOf(arr)
		fmt.Printf("[%d %d %d]\n", val.Index(0).Int(),
			val.Index(1).Int(), val.Index(2).Int()) // [5 6 0]
		fmt.Println(typ.Kind(), typ.String()) // array [3]int

	})

	t.Run("map", func(t *testing.T) {
		var m = map[string]int{ // map
			"tony": 1,
			"jim":  2,
			"john": 3,
		}
		val = reflect.ValueOf(m)
		typ = reflect.TypeOf(m)
		iter := val.MapRange()
		fmt.Printf("{")
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			fmt.Printf("%s:%d,", k.String(), v.Int())
		}
		fmt.Printf("}\n")                     // {tony:1,jim:2,john:3,}
		fmt.Println(typ.Kind(), typ.String()) // map map[string]int
	})

	t.Run("struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		var p = Person{"tony", 23} // 结构体
		val = reflect.ValueOf(p)
		typ = reflect.TypeOf(p)
		fmt.Printf("{%s, %d}\n", val.Field(0).String(),
			val.Field(1).Int()) // {"tony", 23}

		fmt.Println(typ.Kind(), typ.Name(), typ.String()) // struct Person main.Person
	})

	t.Run("channel", func(t *testing.T) {
		var ch = make(chan int, 1) // channel
		val = reflect.ValueOf(ch)
		typ = reflect.TypeOf(ch)
		ch <- 17
		v, ok := val.TryRecv()
		if ok {
			fmt.Println(v.Int()) // 17
		}
		fmt.Println(typ.Kind(), typ.String()) // chan chan int
	})
	t.Run("custom", func(t *testing.T) {
		// 其他自定义类型
		type MyInt int

		var mi MyInt = 19
		val = reflect.ValueOf(mi)
		typ = reflect.TypeOf(mi)
		fmt.Println(typ.Name(), typ.Kind(), typ.String(), val.Int()) // MyInt int main.MyInt 19
	})
}
