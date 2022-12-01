package math_test

import (
	"math/big"

	"github.com/duke-git/lancet/v2/lancetconstraints"
)

func Square(val int64) *big.Int {
	in := big.NewInt(val)
	return in.Mul(in, in)
}

// 泛型版本的Max
func Max[T lancetconstraints.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T lancetconstraints.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
