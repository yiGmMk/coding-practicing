package main

import (
	"fmt"
	"unsafe"
)

/*
使用uintptr类型变量保存栈上变量的地址同样是有风险的，
因为Go使用的是连续栈的栈管理方案，每个goroutine的默认栈大小为2KB（_StackMin = 2048）。
当goroutine当前剩余栈空间无法满足函数/方法调用对栈空间的需求时，Go运行时就会新分配一块更大的内存空间作为该goroutine的新栈空间，
并将该goroutine的原有栈整体复制过来，这样原栈上分配的变量的地址就会发生变化。
*/
/*
go run -gcflags="-l" base/unsafe/uintptr_stack/main.go                                           (base)
变量x的值=[1 2 3 4 5 6 7 8 9 0]
变量x的地址= 0xc000088ec0
栈扩容后，变量x的地址= 0xc0000b5ec0
栈扩容后，变量x的值=[11 12 13 14 15 16 17 18 19 10]
变量p(uintptr)存储的地址上的值=[1 2 3 4 5 6 7 8 9 0]
变量q(unsafe.Pointer)引用的地址上的值=[11 12 13 14 15 16 17 18 19 10]
*/
func main() {
	var x = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("变量x的值=%d\n", x)
	println("变量x的地址=", &x)

	var p = uintptr(unsafe.Pointer(&x))
	var q = unsafe.Pointer(&x)

	a(x) // 执行一系列函数调用

	// 变更数组x中元素的值
	for i := 0; i < 10; i++ {
		x[i] = x[i] + 10
	}

	println("栈扩容后，变量x的地址=", &x)
	fmt.Printf("栈扩容后，变量x的值=%d\n", x)

	fmt.Printf("变量p(uintptr)存储的地址上的值=%d\n", *(*[10]int)(unsafe.Pointer(p)))
	fmt.Printf("变量q(unsafe.Pointer)引用的地址上的值=%d\n", *(*[10]int)(q))
}

func a(x [10]int) {
	var y [100]int
	b(y)
}

func b(x [100]int) {
	var y [1000]int
	c(y)
}

func c(x [1000]int) {
}
