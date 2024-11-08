package base

import (
	"log"
	"testing"
)

/* defer 求值的时机 f1,f2,f3不同在哪里
 */
func f1(nums []int) {
	for _, v := range nums {
		defer println(v)
	}
}

/*
3
3
3
*/
// func f2(nums []int) {
// 	for _, v := range nums {
// 		defer func() {
// 			println(v)
// 		}()
// 	}
// }

func f3(nums []int) {
	for _, v := range nums {
		defer func(v int) {
			println(v)
		}(v)
	}
}

func TestDefer(t *testing.T) {
	nums := []int{1, 2, 3}
	log.Println("f1")
	f1(nums)

	log.Println("f2")
	// f2(nums)

	log.Println("f3")
	f3(nums)
}

func TestDeferslice(t *testing.T) {
	deferSlice()
	deferSlice2()
}
