package algorithms

import (
	"log"
	"math/rand"
	"testing"

	"github.com/RoaringBitmap/roaring/roaring64"
	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
)

/*
references:
1. roaring https://www.roaringbitmap.org/
2. zhihu   https://zhuanlan.zhihu.com/p/445396980
3. code    https://github.com/RoaringBitmap/roaring
*/

func TestRoaringBitmap(t *testing.T) {
	max := int64(1_0000_0000)
	notIn := datastructure.NewSet[int64]()
	b := roaring64.NewBitmap()
	for i := int64(0); i < max; i++ {
		if rand.Int()%555 == 3 {
			notIn.Add(i)
			continue
		}
		b.Add(uint64(i))
	}

	for i := int64(0); i < max; i++ {
		if !b.Contains(uint64(i)) && !notIn.Contain(i) {
			t.Errorf("%d should be member of Bitmap", i)
		}
	}
	log.Println("ok")
}

func BenchmarkRoaringBitmap3(b *testing.B) {
	max := int64(5_0000_0000)
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 10_000, max}

	bitmap := roaring64.NewBitmap()
	for _, v := range nums {
		bitmap.Add(uint64(v))
	}
	for i := 0; i < b.N; i++ {
		for _, v := range nums {
			if !bitmap.Contains(uint64(v)) {
				bitmap.Add(uint64(v))
			}
		}
	}
}
