package base

import (
	"fmt"
	"testing"
)

// 过多的内存占用
func TestSlice(t *testing.T) {
	t.Run("memory", func(t *testing.T) {
		allocSlice := func(min, high int) []int {
			var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 99: 100}
			fmt.Printf("slice b:len(%d),cap(%d)\n", len(b), cap(b))
			return b[min:high]
		}

		a := allocSlice(3, 7)
		fmt.Printf("slice a:len(%d),cap(%d)\n", len(a), cap(a))
	})

	t.Run("reduce mem usage", func(t *testing.T) {
		allocSlice := func(min, high int) []int {
			var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 99: 100}
			fmt.Printf("slice b: len(%d), cap(%d)\n",
				len(b), cap(b))
			nb := make([]int, high-min, high-min)
			copy(nb, b[min:high])
			return nb
		}
		a := allocSlice(3, 7)
		fmt.Printf("slice a:len(%d),cap(%d)\n", len(a), cap(a))
	})
}

// 数据篡改
func TestSliceDataManipulation(t *testing.T) {
	t.Run("data manipulation", func(t *testing.T) {
		allocSlice := func(min, high int) []int {
			var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			fmt.Printf("slice b:len(%d),cap(%d)\n", len(b), cap(b))
			return b[min:high]
		}

		a1 := allocSlice(3, 7) // cap 7
		fmt.Printf("slice a1:len(%d),cap(%d),elements(%v)\n", len(a1), cap(a1), a1)

		a2 := a1[:6] // 取前6个不会越界
		fmt.Printf("slice a2:len(%d),cap(%d),elements(%v)\n", len(a2), cap(a2), a2)
	})

	t.Run("fix data manipulation", func(t *testing.T) {
		allocSlice := func(min, high int) []int {
			var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			fmt.Printf("slice b:len(%d),cap(%d)\n", len(b), cap(b))

			nb := make([]int, high-min, high-min)
			copy(nb, b[min:high])
			return nb
		}

		a1 := allocSlice(3, 7) // cap 7
		fmt.Printf("slice a1:len(%d),cap(%d),elements(%v)\n", len(a1), cap(a1), a1)

		//a2 := a1[:6] // 取前6越界=>panic
		//fmt.Printf("slice a2:len(%d),cap(%d),elements(%v)\n", len(a2), cap(a2), a2)
	})
}
