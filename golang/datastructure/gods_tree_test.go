package datastructure

import (
	"fmt"
	"log"
	"testing"

	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

func TestRedblackTree(t *testing.T) {
	s := utils.ToString(struct{ L, R int }{L: 1, R: 2})
	log.Println(s)

	tree := redblacktree.NewWith(utils.IntComparator)
	tree.Put(1, 1)
	tree.Put(2, 2)

	fmt.Println(tree)

	tree.Put(3, 3)
	tree.Put(4, 4)
	tree.Put(5, 5)
	tree.Put(6, 6)

	fmt.Println(tree)
}
