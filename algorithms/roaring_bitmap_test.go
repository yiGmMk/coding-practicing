package algorithms

import (
	"testing"

	"github.com/RoaringBitmap/roaring/roaring64"
)

/*
references:
1. roaring https://www.roaringbitmap.org/
2. zhihu   https://zhuanlan.zhihu.com/p/445396980
3. code    https://github.com/RoaringBitmap/roaring
*/

func TestRoaringBitmap(t *testing.T) {

}

func BenchmarkRoaringBitmap3(b *testing.B) {
	max := int64(50_0000_0000)
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
