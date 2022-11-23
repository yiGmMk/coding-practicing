package base

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	size := 64
	b := make([]byte, size)
	n, err := rand.Read(b)
	if err != nil {
		t.Error(err)
	}
	if n != size {
		t.Errorf("size not %d", n)
	}
	fmt.Printf("%x", b)
}
