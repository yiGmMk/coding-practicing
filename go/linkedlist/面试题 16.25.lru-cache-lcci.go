/*
 * @lc app=leetcode.cn id=面试题 16.25 lang=golang
 * @lcpr version=30204
 *
 * [面试题 16.25] LRU 缓存
 */
package linkedlist

import "container/list"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type (
	node struct {
		key, val int
		next     *node
	}
	LRUCache struct {
		num, capacity int
		kv            map[int]*list.Element
		l             *list.List
	}
)

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		kv:       make(map[int]*list.Element, capacity),
		l:        list.New(),
	}
}

// swap by lru
func (this *LRUCache) Get(key int) int {
	if v, ok := this.kv[key]; ok {
		this.l.Remove(v)
		this.kv[key] = this.l.PushFront(v.Value)
		return v.Value.(node).val
	}
	// 未找到
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.kv[key]; ok {
		this.l.Remove(v)
	} else {
		this.num++
	}
	this.kv[key] = this.l.PushFront(node{
		key: key,
		val: value,
	})
	if this.num > this.capacity {
		b := this.l.Back()
		delete(this.kv, b.Value.(node).key)
		this.l.Remove(b)
		this.num--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
