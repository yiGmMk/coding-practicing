package algorithms

import (
	"github.com/yiGmMk/leetcode/algorithms"
)

// return -1 of array if not found the target,result is different with sort.SearchInts
// 没有搜索到目标返回-1
// 与sort.SearchInts()行为不同,sort.SearchInts()返回元素插入位置,可能是len(values)
func BinarySearch[T algorithms.Ordered](values []T, target T) int {
	size := len(values)
	left := 0
	right := size - 1
	for left <= right {
		mid := left + (right-left)/2
		if values[mid] == target {
			return mid
		} else if values[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
