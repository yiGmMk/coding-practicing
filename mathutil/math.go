package mathutil

import (
	"github.com/duke-git/lancet/v2/lancetconstraints"
)

func Min[T lancetconstraints.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
