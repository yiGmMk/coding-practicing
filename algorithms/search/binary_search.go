package algorithms

import (
	"github.com/duke-git/lancet/v2/lancetconstraints"
)

// return size of array if not found the target
// 没有搜索到目标返回数组长度
func BinarySearch[T lancetconstraints.Number](values []T, target T) int {
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
