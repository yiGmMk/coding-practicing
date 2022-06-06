package jzoffer

import "github.com/emirpasic/gods/trees/redblacktree"

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
