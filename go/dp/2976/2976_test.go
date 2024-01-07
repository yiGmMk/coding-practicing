package jzoffer

import "testing"

func Test2976(t *testing.T) {
	// 输入：source = "aaaa", target = "bbbb", original = ["a","c"], changed = ["c","b"], cost = [1,2]
	// 输出：12
	if got := minimumCost("aaaa", "bbbb",
		[]byte{'a', 'c'},
		[]byte{'c', 'b'}, []int{1, 2}); got != 12 {
		t.Errorf("minimumCost() = %v, want 12", got)
	}

	// 输入：source = "abcd", target = "acbe",
	// original = ["a","b","c","c","e","d"], changed = ["b","c","b","e","b","e"],
	// cost = [2,5,5,1,2,20]
	// 输出：28
	if got := minimumCost("abcd", "acbe",
		[]byte("a,b,c,c,e,d"),
		[]byte("b,c,b,e,b,e"), []int{2, 5, 5, 1, 2, 20}); got != 28 {
		t.Errorf("minimumCost() = %v, want 28", got)
	}
}
