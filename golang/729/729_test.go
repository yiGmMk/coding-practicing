package jzoffer

import (
	"testing"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func TestCases(t *testing.T) {
	type inNode struct {
		start, end int
	}
	ts := []struct {
		nodes []struct {
			val inNode
			res bool
		}
	}{
		// 测试用例1
		{
			nodes: []struct {
				val inNode
				res bool
			}{
				{inNode{start: 10, end: 20}, true},
				{inNode{start: 15, end: 25}, false},
				{inNode{start: 20, end: 30}, true},
				{inNode{start: 5, end: 15}, false},
			},
		},
		// 测试用例2
		{
			// [47,50],[33,41],[39,45],[33,42],[25,32],[26,35],[19,25],[3,8],[8,13],[18,27]]
			// [true,    true,   false, false,  true,   false,   true,   true, true, false]
			nodes: []struct {
				val inNode
				res bool
			}{
				{inNode{start: 47, end: 50}, true},
				{inNode{start: 33, end: 41}, true},
				{inNode{start: 39, end: 45}, false},
				{inNode{start: 33, end: 42}, false},
				{inNode{start: 25, end: 32}, true},
				{inNode{start: 26, end: 35}, false},
				{inNode{start: 19, end: 25}, true},
				{inNode{start: 3, end: 8}, true},
				{inNode{start: 8, end: 13}, true},
				{inNode{start: 18, end: 27}, false},
			},
		},
	}
	var handdler Handler
	for _, tt := range ts {
		handdler = Constructor1()
		for _, n := range tt.nodes {
			if handdler.Book(n.val.start, n.val.end) != n.res {
				t.Errorf("error,node:%+v", n)
			}
		}
	}
}

type Handler interface {
	Book(start, end int) bool
}

// -----------------------1. 红黑树实现--------------------------
// 作者：endlesscheng
// 链接：https://leetcode.cn/problems/my-calendar-i/solution/go-ping-heng-shu-by-endlesscheng-ihq4/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
type MyCalendar1 struct {
	*redblacktree.Tree
}

func Constructor1() *MyCalendar1 {
	t := redblacktree.NewWithIntComparator()
	t.Put(-1, -1) // 哨兵
	return &MyCalendar1{t}
}

func (c *MyCalendar1) Book(start, end int) bool {
	floor, _ := c.Floor(start)
	if floor.Value.(int) > start { // [start,end) 左侧区间的右端点超过了 start
		return false
	}
	if it := c.IteratorAt(floor); it.Next() && it.Key().(int) < end { // [start,end) 右侧区间的左端点小于 end
		return false
	}
	c.Put(start, end) // 可以插入区间 [start,end)
	return true
}
