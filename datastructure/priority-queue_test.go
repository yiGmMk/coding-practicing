package datastructure

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	pq := PQ{}
	heap.Init(&pq)

	heap.Push(&pq, &entry{key: "1", priority: 1})
	heap.Push(&pq, &entry{key: "2", priority: 2})
	heap.Push(&pq, &entry{key: "3", priority: 3})
	heap.Push(&pq, &entry{key: "40", priority: 4})
	heap.Push(&pq, &entry{key: "55", priority: 5})

	for len(pq) > 0 {
		fmt.Println(heap.Pop(&pq))
	}
}
