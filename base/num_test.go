package base

import (
	"log"
	"math"
	"testing"
)

func TestUint(t *testing.T) {
	num := uint32(0)
	max, maxAdd := uint32(math.MaxUint32), math.MaxUint32+1
	log.Println(num-1, max, maxAdd)

	intMax := int32(math.MaxInt32)
	log.Println(math.MinInt32, math.MaxInt32, intMax+1)
}
