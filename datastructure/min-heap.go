package datastructure

type maxHeap []int

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *maxHeap) Push(v any) {
	*h = append(*h, v.(int))
}

// 小顶堆
type minHeap []int

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *minHeap) Push(v any) {
	*h = append(*h, v.(int))
}
