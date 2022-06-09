package base

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestCap(t *testing.T) {
	// 切片的大小和空间分配
	array := make([]int, 0, 1)
	fmt.Printf("array,size:%d,p_size:%d,len:%d,capacity:%d\n",
		unsafe.Sizeof(array), unsafe.Sizeof(&array), len(array), cap(array))
}
