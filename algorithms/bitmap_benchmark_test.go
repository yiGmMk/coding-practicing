package algorithms

import "testing"

func BenchmarkBitmap1(b *testing.B) {
	max := int64(10_0000)
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 10_000, max}
	bitmap := NewBitmap(max)
	for _, v := range nums {
		bitmap.Set(v)
	}
	for i := 0; i < b.N; i++ {
		for _, v := range nums {
			if !bitmap.Get(v) {
				bitmap.Set(v)
			}
		}
	}
}

func BenchmarkBitmap2(b *testing.B) {
	max := int64(10_0000_0000)
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 10_000, max}
	bitmap := NewBitmap(max)
	for _, v := range nums {
		bitmap.Set(v)
	}
	for i := 0; i < b.N; i++ {
		for _, v := range nums {
			if !bitmap.Get(v) {
				bitmap.Set(v)
			}
		}
	}
}

func BenchmarkBitmap3(b *testing.B) {
	max := int64(50_0000_0000)
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 10_000, max}
	bitmap := NewBitmap(max)
	for _, v := range nums {
		bitmap.Set(v)
	}
	for i := 0; i < b.N; i++ {
		for _, v := range nums {
			if !bitmap.Get(v) {
				bitmap.Set(v)
			}
		}
	}
}
