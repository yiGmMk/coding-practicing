package math_test

import "math/big"

func Square(val int64) *big.Int {
	in := big.NewInt(val)
	return in.Mul(in, in)
}
