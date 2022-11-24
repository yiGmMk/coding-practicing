package base

import (
	"fmt"
	"testing"
	"unsafe"
)

// chapter9/sources/go-unsafe/funcs_in_unsafe.go
type Foo struct {
	a int
	b string
	c [10]byte
	d float64
}

var i int = 5
var a = [100]int{}
var sl = a[:]
var f Foo

func TestSizeof(t *testing.T) {
	fmt.Println(unsafe.Sizeof(i))  // 8
	fmt.Println(unsafe.Sizeof(a))  // 800
	fmt.Println(unsafe.Sizeof(sl)) // 24 (注：返回的是切片描述符的大小)
	fmt.Println(unsafe.Sizeof(f))  // 48=8+16+10+(6,padding)+8
	fmt.Println(
		unsafe.Sizeof(f.a), //8
		unsafe.Sizeof(f.b), //16
		unsafe.Sizeof(f.c), //10
		unsafe.Sizeof(f.d)) //8
	fmt.Println(unsafe.Sizeof((*int)(nil))) // 8
}

func TestAlign(t *testing.T) {
	fmt.Println(unsafe.Alignof(i))          // 8
	fmt.Println(unsafe.Alignof(f.a))        // 8
	fmt.Println(unsafe.Alignof(a))          // 8
	fmt.Println(unsafe.Alignof(sl))         // 8
	fmt.Println(unsafe.Alignof(f))          // 8
	fmt.Println(unsafe.Alignof(f.c))        // 1
	fmt.Println(unsafe.Alignof(struct{}{})) // 1 (注：空结构体的对齐系数为1)
	fmt.Println(unsafe.Alignof([0]int{}))   // 8 (注：长度为0的数组，其对齐系数依然与其元素
	// 类型的对齐系数相同)
}

func TestOffsetof(t *testing.T) {
	fmt.Println(unsafe.Offsetof(f.b)) // 8
	fmt.Println(unsafe.Offsetof(f.d)) // 40
}

func TestPointer(t *testing.T) {
	// 利用Pointer可以突破go的类型安全保护
	var a uint32 = 0x12345678
	fmt.Printf("0x%x\n", a) // 0x12345678

	p := (unsafe.Pointer)(&a) // 利用unsafe.Pointer的性质1

	b := (*[4]byte)(p) // 利用unsafe.Pointer的性质2
	b[0] = 0x23
	b[1] = 0x45
	b[2] = 0x67
	b[3] = 0x8a

	fmt.Printf("0x%x\n", a) // 0x8a674523 (注：在小端字节序系统中输出此值)
}
