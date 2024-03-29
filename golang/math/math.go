package math_test

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

func Square(val int64) *big.Int {
	in := big.NewInt(val)
	return in.Mul(in, in)
}

// 泛型版本的Max
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
