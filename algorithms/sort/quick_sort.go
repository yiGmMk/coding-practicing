package algorithms

import (
	"github.com/yiGmMk/leetcode/algorithms"
)

/*快排在partition时,选择pivot有多种选择:最右边的,随机一个,受数据本身的有序性影响,最好时是n*log(n)
 */
func QuickSort[T algorithms.Ordered](values []T) {
	quickSort(values, 0, len(values)-1)
}

//partition returns the index of the pivot after partitioning the array.
func partition[T algorithms.Ordered](values []T, left, right int) int {
	pivot := values[right]
	index := left - 1
	for i := left; i < right; i++ {
		if values[i] <= pivot { // 找到小于/等于pivot的数与大于pivot的数交换位置
			index++
			values[index], values[i] = values[i], values[index]
		}
	}
	values[index+1], values[right] = values[right], values[index+1]
	return index + 1
}

func quickSort[T algorithms.Ordered](values []T, left, right int) {
	if len(values) <= 1 {
		return
	}

	if left < right {
		pivot := partition(values, left, right)
		quickSort(values, left, pivot-1)
		quickSort(values, pivot+1, right)
	}
}
