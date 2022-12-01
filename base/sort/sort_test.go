package base

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"

	"github.com/yiGmMk/leetcode/datastructure"
)

func pivotInteger(n int) int {
	// m[i]=m[n]-m[i-1]
	sum := make([]int, n+1)
	s := 0
	for i := 0; i <= n; i++ {
		sum[i] = s + i
		s = s + i
	}

	//8,6 36-21+6

	for i, v := range sum {
		if v == sum[n]-v+i {
			return i
		}
	}

	return -1
}

func Test6245(t *testing.T) {
	fmt.Println(pivotInteger(8))
}

func appendCharacters(s string, t string) int {
	last := 0
	add := 0
	for i, v := range t {
		if last < len(s) && strings.IndexRune(s[last:], v) >= last {
			continue
		}
		add++
		if i != len(t)-1 {
			last = len(s)
		}
		last = i
	}
	return add
}

type ListNode = datastructure.ListNode

// func removeNodes(head *ListNode) *ListNode {
// 	type bnode struct {
// 		big  int
// 		val  int
// 		node *ListNode
// 	}
// 	ns := []*bnode{}
// 	max := math.MinInt64
// 	for p := head; p != nil; p = p.Next {
// 		if p.Val > max {
// 			max = p.Val
// 			if p != head {
// 				ns = []*bnode{}
// 			}
// 			max
// 		}
// 		ns = append(ns, &bnode{big: max, val: p.Val, node: p})
// 	}

// }

func Test6247(t *testing.T) {

}

func Convey(t *testing.T, got, expect int) {
	if got != expect {
		t.Errorf("got:%d,wangt:%d", got, expect)
	}
}

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
	if found != true {
		t.Errorf("find %d failed", len(nums))
	}

	id := sort.SearchInts(nums, 11)
	Convey(t, id, len(nums))

	id = sort.SearchInts(nums, 10)
	Convey(t, id, len(nums)-1)
}
