package concurrency

import (
	"fmt"
	"log"
	"testing"

	set "github.com/duke-git/lancet/v2/datastructure/set"
	tree "github.com/duke-git/lancet/v2/datastructure/tree"
	"github.com/duke-git/lancet/v2/strutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLancetStrUtil(t *testing.T) {
	s := "hello"
	rs := strutil.ReverseStr(s)
	fmt.Println(rs) //olleh
}

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}
func TestTreeNode(t *testing.T) {
	Convey("tree", t, func() {
		tt := tree.NewBSTree[int](1, &intComparator{})
		tt.InsertNode(2)
		tt.InsertNode(3)
		tt.InsertNode(4)
		tt.InsertNode(2)
		tt.InsertNode(3)
		tt.InsertNode(4)
		tt.InsertNode(-2)
		tt.Print()
	})
}

func TestSet(t *testing.T) {
	Convey("set", t, func() {
		s := set.NewSet[int](1, 2, 3)
		s.Add(4, 5, 6)
		s.Add(4, 5, 6)
		s.Add(7, 8, 9)

		size := s.Size()
		log.Println("size:", size)
		So(size, ShouldEqual, 9)

		s.Iterate(func(val int) {
			log.Println("val:", val)
		})

		exist := s.Contain(1)
		So(exist, ShouldEqual, true)

		exist = s.ContainAll(set.NewSet[int](1, 7, 9))
		So(exist, ShouldEqual, true)

		s.Delete(9)
		exist = s.Contain(9)
		So(exist, ShouldEqual, false)

		intersecter := s.Intersection(set.NewSet[int](1, 2, 7, 8, 9))
		equal := intersecter.Equal(set.NewSet[int](1, 2, 7, 8))
		So(equal, ShouldEqual, true)

		minus := s.Minus(set.NewSet[int](1, 2, 3, 44))
		log.Println("minus:", minus)

		diff := s.SymmetricDifference(set.NewSet[int](1, 2, 3, 44))
		log.Println("diff:", diff)

		vs := s.Values()
		log.Println("values:", vs)
	})
}
