package algorithms

import (
	"fmt"

	"github.com/yiGmMk/leetcode/algorithms"
)

var debug = false

func BubbleSort[T algorithms.Ordered](values []T) {
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values)-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
		if debug {
			fmt.Println(values)
		}
	}
}
