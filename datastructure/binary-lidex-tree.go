package datastructure

// TODO:
// 树状数组,支持:单点修改,区间查询
// 1. https://oi-wiki.org/ds/fenwick/
// 2. https://zhuanlan.zhihu.com/p/93795692
type BinaryIndexedTree struct {
	tree []int
	cap  int
}

// lowbit
