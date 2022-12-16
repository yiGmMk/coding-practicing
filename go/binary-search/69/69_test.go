package jzoffer

import (
	"math"
	"testing"
)

func Test69(t *testing.T) {
	for i := 0; i < 101; i++ {
		got := int64(mySqrt(i))
		v := int64(math.Sqrt(float64(i)))
		if v != got {
			t.Errorf("sqrt of %d.got:%d,not:%d", i, got, v)
		}
	}
}
