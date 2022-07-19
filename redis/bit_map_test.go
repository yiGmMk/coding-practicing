package redis

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	r := NewRedisInmemory(t)
	bitKey := "bit_map_test"
	count, err := r.BitCount(bitKey, 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatal("count should be 0")
	}
	for i := int64(0); i < 10000; /*_0000*/ i++ {
		r.SetBit(bitKey, i, 1)
	}
	count, err = r.BitCount(bitKey, 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(count)
}
