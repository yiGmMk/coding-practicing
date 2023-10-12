package base

import (
	"fmt"
	"log"
	"testing"
	"unsafe"
)

func TestCap(t *testing.T) {
	// 切片的大小和空间分配
	array := make([]int, 0, 1)
	fmt.Printf("array,size:%d,p_size:%d,len:%d,capacity:%d\n",
		unsafe.Sizeof(array), unsafe.Sizeof(&array), len(array), cap(array))
}

// 扩容
func TestAppend(t *testing.T) {
	arr := make([]int, 0)
	last := 0
	for i := 0; i <= 2000; i++ {
		if cap(arr) != last {
			fmt.Print(cap(arr), " ")
			last = cap(arr)
		}
		arr = append(arr, i)
	}

	// go1.13=> 	1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.14=> 	1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.15=> 	1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.16=> 	1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.17=> 	1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.17.9 =>  1 2 4 8 16 32 64 128 256 512 1024 1280 1696 2304
	// go1.18.10=>  1 2 4 8 16 32 64 128 256 512 848 1280 1792 2560
	// go1.19=> 	1 2 4 8 16 32 64 128 256 512 848 1280 1792 2560
	// go1.20=>		1 2 4 8 16 32 64 128 256 512 848 1280 1792 2560
	// go1.21=> 	1 2 4 8 16 32 64 128 256 512 848 1280 1792 2560

	a := []int{1, 2}
	fmt.Println("\nlen=", len(a), "cap=", cap(a))
	a = append(a, 1, 2, 3)
	fmt.Println("len=", len(a), "cap=", cap(a))
}

func TestSliceStorge(t *testing.T) {
	// reslicing创建的新切片append是什么行为,会产生什么结果
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(len(s), cap(s), "s:", s)

	s1 := s[2:5]
	log.Println(len(s1), cap(s1), "s1:", s1)
	s1 = append(s1, 100)
	log.Println(len(s1), cap(s1), "s1:", s1)
	log.Println(len(s), cap(s), "s:", s)

	/*s1通过reslicing创建,和s共享一个底层数组,append后修改了共同的数组值
	 *2022/07/31 15:27:37 10 10 s:  [0 1 2 3 4 5 6 7 8 9]
	 *2022/07/31 15:27:38 3 8   s1: [2 3 4]
	 *2022/07/31 15:27:38 4 8   s1: [2 3 4 100]
	 *2022/07/31 15:29:19 10 10 s:  [0 1 2 3 4 100 6 7 8 9]
	 */

	// append到原有的底层数组不够装的时候,会与原数组脱离,s1的底层数组不再是s的底层数组
	s1 = append(s1, 200, 300, 400, 500, 600, 700, 800, 900, 1000)
	log.Println(len(s1), cap(s1), "s1:", s1)
	log.Println(len(s), cap(s), "s:", s)
	/*
	 *2022/07/31 16:09:29 13 16 s1: [2 3 4 100 200 300 400 500 600 700 800 900 1000]
	 *2022/07/31 16:09:49 10 10 s:  [0 1 2 3 4 100 6 7 8 9]
	 */
}
