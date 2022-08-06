package base

import (
	"log"
	"sort"
	"testing"
)

func TestFind(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 5, 5, 5, 6, 7, 8, 9, 10}
	target := 6
	i, found := sort.Find(len(nums), func(n int) int {
		if nums[n] == target {
			return 0
		}
		if nums[n] > target {
			return -1
		}
		return 1
	})
	log.Println(i, found)
}
