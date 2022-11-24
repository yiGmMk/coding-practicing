// chapter9/sources/go-unsafe/go_mem_obj_ref_unsafepointer_vs_uintptr.go
package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

type Foo struct {
	name string
}

func finalizer(p *Foo) {
	fmt.Printf("Foo: [%s]被垃圾回收\n", p.name)
}

func NewFoo(name string) *Foo {
	var f Foo = Foo{
		name: name,
	}
	runtime.SetFinalizer(&f, finalizer)
	return &f
}

func allocLargeObject() *[1000000]uint64 {
	a := [1000000]uint64{}
	return &a
}

/*
unsafe.Pointer和其他常规类型指针一样，可以作为对象引用。
如果一个对象仍然被某个unsafe.Pointer变量引用着，那么该对象是不会被垃圾回收的。
但是uintptr并不是指针，它仅仅是一个整型值，即便它存储的是某个对象的内存地址，它也不会被算作对该对象的引用。
如果认为将对象地址存储在一个uintptr变量中，该对象就不会被垃圾回收器回收，那就是对uintptr的最大误解。
*/
func main() {
	var p1 = uintptr(unsafe.Pointer(NewFoo("FooRefByUintptr")))
	var p2 = unsafe.Pointer(NewFoo("FooRefByPointer"))

	for i := 0; i < 5; i++ {
		allocLargeObject()

		// 尝试输出p1和p2地址上的值
		q1 := (*Foo)(unsafe.Pointer(p1))
		fmt.Printf("object ref by uintptr: %+v\n", *q1)

		q2 := (*Foo)(p2)
		fmt.Printf("object ref by pointer: %+v\n", *q2)

		runtime.GC() // 运行垃圾回收
		time.Sleep(1 * time.Second)
	}
}
