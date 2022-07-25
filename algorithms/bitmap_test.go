package algorithms

import (
	"log"
	"math/rand"
	"testing"
	"unsafe"

	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
)

func TestBitmap1(t *testing.T) {
	max := (1_0000_0000)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 1000_000, max}
	b := NewBitmap(max)
	for _, v := range nums {
		b.Set(v)
	}

	log.Println(
		"len:", len(b),
		"size of b:", unsafe.Sizeof(b[:]),
		"cap:", cap(b), "Bytes:", cap(b)*8, "MB:", float64(cap(b)*8)/1024/1024,
		"origin MB:", max*8/1024/1024.0)

	for _, v := range nums {
		if !b.Get(v) {
			t.Errorf("%d not in set", v)
		}
	}
}

func TestBitmap2(t *testing.T) {
	max := (1_0000_0000)
	notIn := datastructure.NewSet[int]()
	b := NewBitmap(max)
	for i := (0); i < max; i++ {
		if rand.Int()%555 == 3 {
			notIn.Add(i)
			continue
		}
		b.Set(i)
	}

	log.Println("len:", len(b), "size of b:", unsafe.Sizeof(b))

	for i := (0); i < max; i++ {
		if !b.Get(i) && !notIn.Contain(i) {
			t.Errorf("%d should be member of Bitmap", i)
		}
	}
	log.Println("ok")
}
