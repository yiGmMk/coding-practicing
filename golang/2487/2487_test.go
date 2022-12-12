package jzoffer

import (
	"fmt"
	"testing"

	"github.com/yiGmMk/leetcode/datastructure"
)

var _2487TestCase = []struct {
	in     []int
	expect []int
}{
	{},
}

func Test2487(t *testing.T) {

}

func TestArray(t *testing.T) {
	in := []int{5, 2, 13, 3, 8}
	out := []int{}
	for _, v := range in {
		if len(out) > 0 {
			j, big := len(out)-1, false
			for i := j; i >= 0; i-- {
				if v > out[i] {
					j = i
					big = true
				}
			}
			if big {
				if j == 0 {
					out = []int{}
				} else {
					out = out[:j]
				}
			}
		}
		out = append(out, v)
	}
	fmt.Println(len(out), cap(out), out)
}

func TestList(t *testing.T) {
	in := &datastructure.ListNode{
		Val: 5,
		Next: &datastructure.ListNode{
			Val: 2,
			Next: &datastructure.ListNode{
				Val: 13,
				Next: &datastructure.ListNode{
					Val: 3, Next: &datastructure.ListNode{
						Val: 8, Next: nil}},
			},
		}}
	removeNodes(in)
}
