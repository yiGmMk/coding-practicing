package jzoffer

import (
	"fmt"
	"testing"
)

func Test3(t *testing.T) {

	// "aaadbdcdac"
	// "cdbabaddba"
	// ["a","c","b","d","b","a","c"]
	// ["c","a","d","b","c","b","d"]
	// [ 7,  2,  1,  3,  6,  1,  7]
	fmt.Println(minimumCost("aaadbdcdac", "cdbabaddba",
		[]byte{'a', 'c', 'b', 'd', 'b', 'a', 'c'},
		[]byte{'c', 'a', 'd', 'b', 'c', 'b', 'd'},
		[]int{7, 2, 1, 3, 6, 1, 7}))

	// "aaaabadaaa"
	// "dbdadddbad"
	// ["c","a","c","a","a","b","b","b","d","d","c"]
	// ["a","c","b","d","b","c","a","d","c","b","d"]
	// [7,8,11,9,7,6,4,6,9,5,9]
	fmt.Println(minimumCost("aaaabadaaa", "dbdadddbad",
		[]byte{'c', 'a', 'c', 'a', 'a', 'b', 'b', 'b', 'd', 'd', 'c'},
		[]byte{'a', 'c', 'b', 'd', 'b', 'c', 'a', 'd', 'c', 'b', 'd'},
		[]int{7, 8, 11, 9, 7, 6, 4, 6, 9, 5, 9}))

	// "abcd"
	// "acbe"
	// ["a","b","c","c","e","d"]
	// ["b","c","b","e","b","e"]
	// [2,5,5,1,2,20]
	fmt.Println(minimumCost("abcd", "acbe",
		[]byte{'a', 'b', 'c', 'c', 'e', 'd'},
		[]byte{'b', 'c', 'b', 'e', 'b', 'e'},
		[]int{2, 5, 5, 1, 2, 20}))

	// source = "aaaa", target = "bbbb", original = ["a","c"], changed = ["c","b"], cost = [1,2]
	// 输出：12
	fmt.Println(minimumCost("aaaa", "bbbb",
		[]byte{'a', 'c'}, []byte{'c', 'b'}, []int{1, 2}))

	// source = "abcd", target = "abce", original = ["a"], changed = ["e"], cost = [10000]
	fmt.Println(minimumCost("abcd", "abce",
		[]byte{'a'}, []byte{'e'}, []int{10000}))

}
