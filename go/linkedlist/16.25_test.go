package linkedlist

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	h := Constructor(2)
	fmt.Println(2, ":", h.Get(2))
	h.Put(1, 1)
	h.Put(2, 2)
	fmt.Println(2, ":", h.Get(2))
	h.Put(3, 3)
	fmt.Println(3, ":", h.Get(3))
	fmt.Println(1, ":", h.Get(1))
	fmt.Println(2, ":", h.Get(2))
}

func TestLru2(t *testing.T) {
	h := Constructor(1)
	h.Put(2, 1)
	fmt.Println(2, ":", h.Get(2))
	h.Put(3, 2)
	fmt.Println(3, ":", h.Get(3))
	fmt.Println(2, ":", h.Get(2))
}
