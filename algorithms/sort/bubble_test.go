package algorithms

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	debug = true
	BubbleSort(values)
	fmt.Println("sorted:", values)
}

type bubble2 []int

func (b bubble2) Len() int {
	return len(b)
}

func (b bubble2) swapFirstTwoCards() {
	if b.Len() < 2 {
		return
	}
	b[0], b[1] = b[1], b[0]
}

// 将第一个元素移至最后
func (b bubble2) putFirstCardToEnd() {
	if b.Len() < 2 {
		return
	}
	b0 := b[0]
	for i := 0; i < b.Len()-1; i++ {
		b[i] = b[i+1]
	}
	b[b.Len()-1] = b0
}

func (b bubble2) First() int {

	return b[0]
}

func (b bubble2) Second() int {
	return b[1]
}

func TestOP(t *testing.T) {
	b := bubble2([]int{1, 4, 5, 2, 3})
	b.putFirstCardToEnd()
	fmt.Println(b)
}

// 如何将一副扑克牌排序，限制条件是只能查看最上面的两张牌，交换最上面的两张牌，或是将最上面的一张牌放到这摞牌的最下面。
// https://www.panjinbo.com/blogs/algorithm-card-deck-sorting/
func TestBubbleCards(t *testing.T) {
	b := bubble2([]int{1, 4, 5, 2, 3})
	n := b.Len()
	for i := 0; i < n; i++ {
		// 每次循环找到第i小的元素,循环后 [n-i-1,n-1]是有序的
		for j := 0; j < n-i-1; j++ {
			if b.First() < b.Second() {
				b.swapFirstTwoCards()
			}
			b.putFirstCardToEnd()
		}
		fmt.Println(b)

		// index 0 is the i-th smallest card
		// index 1 to i-1 is the ordered 0 to i-1 the smallest card
		// put [1, i-1]th to the end
		// 将  [最小的数|已排序的数|未排序的数]
		// 变为[最小的数|未排序的数|已排序的数]
		for k := 0; k < i; k++ {
			b.swapFirstTwoCards()
			b.putFirstCardToEnd()
		}
		b.putFirstCardToEnd()
		fmt.Println(b)
	}
	fmt.Println(b)
}

// 排序, 每轮循环 先找出未排序的最小的数
// 变成
// [最小的数|已排序的数|未排序的数]
// 然后排序最小的数
// 一轮循环完成
