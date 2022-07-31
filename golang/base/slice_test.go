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

func TestSlice(t *testing.T) {
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
}
