package jzoffer

import (
	"fmt"
	"testing"
)

func Test187(t *testing.T) {
	for _, v := range []string{"AAAAAAAAAAA", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"} {
		got := findRepeatedDnaSequences(v)

		fmt.Println(v, "got=", got)
	}
}
