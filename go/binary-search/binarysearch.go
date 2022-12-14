package binarysearch

import (
	"sort"

	"golang.org/x/exp/constraints"
)

/*
	func SearchInts(a []int, x int) int {
		return Search(len(a), func(i int) bool { return a[i] >= x })
	}
*/
func LowerBound[T constraints.Ordered](v []T, target T) int {
	return sort.Search(len(v), func(i int) bool { return v[i] >= target })
}

func UpperBound[T constraints.Signed | constraints.Float](v []T, target T) int {
	return sort.Search(len(v), func(i int) bool { return v[i] >= (target + 1) })
}

func UpperBound1[T constraints.Signed | constraints.Float](v []T, target T) int {
	return sort.Search(len(v), func(i int) bool { return v[i] > target })
}
