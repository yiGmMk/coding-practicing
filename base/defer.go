package base

import "fmt"

// defer 求值的时机,slice
/*
 * 1. [3/3]0xc0000e4018 3 [1 2 3]
 *
 * 1. [1/1]0xc0000b0458 1 [3]
 *
 * 1.in defer,nums: [1/1]0xc0000b0458 1 [3]
 *
 * 1.in defer,n: [3/3]0xc0000e4018 3 [1 2 3]
 *
 * 2. [1/1]0xc0000b4490 1 [3]
 *
 * 2.in defer [1/1]0xc0000b4490 1 [3]
 */
func deferSlice() {
	nums := []int{1, 2, 3}
	println("1.", nums, len(nums), fmt.Sprintln(nums))

	defer func(n []int) {
		println("1.in defer,n:", n, len(n), fmt.Sprintln(n))
		// 切片不是公用底层数组吗,为什么这里不是打印3
	}(nums)

	defer func() {
		println("1.in defer,nums:", nums, len(nums), fmt.Sprintln(nums))
		// 切片不是公用底层数组吗,为什么这里不是打印3
	}()

	nums = []int{3}
	_ = nums
	println("1.", nums, len(nums), fmt.Sprintln(nums))
	// 两边地址不同了,这是为什么?

	/// nums = []int{3},重新赋值后,nums指向的底层数据变量,不是初始的那个,defer里的仍执行原始的那个
}

func deferSlice2() {
	nums := []int{1, 2, 3}

	defer func(n *[]int) {
		println("2.in defer", *n, len(*n), fmt.Sprintln(*n))
	}(&nums)

	nums = []int{3}
	_ = nums
	println("2.", nums, len(nums), fmt.Sprintln(nums))
}
