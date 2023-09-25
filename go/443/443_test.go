package jzoffer

import (
	"fmt"
	"testing"
)

func Test443(t *testing.T) {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}
