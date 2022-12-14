package binarysearch

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

var bsIns = []struct {
	v          []int
	target     int
	lowerbound int
	upperbound int
}{
	{
		v:          []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:     11,
		lowerbound: 11,
		upperbound: 11,
	},
	{
		v:          []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:     -1,
		lowerbound: 0,
		upperbound: 0,
	},
	{
		v:          []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		target:     4,
		lowerbound: 4,
		upperbound: 5,
	},
}

func TestBinarySearch(t *testing.T) {
	for i, v := range bsIns {
		got := LowerBound(v.v, v.target)
		if got != v.lowerbound {
			t.Errorf("[%d].lowerbound.got:%d,want:%d", i, got, v.lowerbound)
		}

		got = UpperBound(v.v, v.target)
		if got != v.upperbound {
			t.Errorf("[%d].upperbound.got:%d,want:%d", i, got, v.lowerbound)
		}
	}
}

func TestSearch(t *testing.T) {
	//求 int(sqrt(90))
	//x*x<=90 => 先 true 后 false 的 f(x)
	x := sort.Search(12, func(x int) bool {
		x++
		return x*x > 90
	})
	fmt.Printf("x*x>90,x=%d\n", x)

	i := strings.Index("a gopher", "gopher")
	iCode := strings.Index("a gopher", "code")
	fmt.Printf("index of `gopher`:%d\nindex of `code`  :%d\n", i, iCode)
}
