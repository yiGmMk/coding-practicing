package datastructure

import "container/heap"

// PriorityQueue :https://www.liaoxuefeng.com/wiki/1252599548343744/1265120632401152
// Queue是先进先出(FIFO)的队列
// PriorityQueue与Queue的区别在于,出队顺序与元素的优先级有关,优先出优先级最高的元素
// PriorityQueue可以使用数组/队列/堆等实现,使用堆是综合效率最高

// https://github.com/halfrost/LeetCode-Go/blob/master/structures/PriorityQueue.go
// entry 是 priorityQueue 中的元素
type entry struct {
	key      string
	priority int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp, ok := x.(*entry)
	if ok {
		temp.index = len(*pq)
		*pq = append(*pq, temp)
	}
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// update modifies the priority and value of an entry in the queue.
func (pq *PQ) update(entry *entry, value string, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
