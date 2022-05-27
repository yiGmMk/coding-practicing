package math_test

import (
	"math"
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"
)

func FuzzBig(f *testing.F) {
	ts := []int64{math.MaxInt64, 0, 1, 2, 3, 4, 1e10, 1111_2222_3333_4444}
	for i := 10000; i > 0; i-- {
		ts = append(ts, rand.Int63())
	}
	for _, t := range ts {
		f.Add(t)
	}
	f.Fuzz(func(t *testing.T, orig int64) {
		res := Square(orig)
		if orig == 0 {
			cmp := res.Cmp(big.NewInt(0))
			if cmp != 0 {
				t.Errorf("calc error,cmp:%d", res.Cmp(big.NewInt(0)))
			}
			return
		}
		if res.Div(res, big.NewInt(orig)).Cmp(big.NewInt(orig)) != 0 {
			t.Errorf("calc error")
		}
	})
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := strings.Replace(orig, "e", "a", -1)
		doubleRev := strings.Replace(orig, "a", "e", -1)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
