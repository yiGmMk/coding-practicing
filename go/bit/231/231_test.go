package jzoffer

import (
	"fmt"
	"strconv"
	"testing"
)

func Test231(t *testing.T) {
	fmt.Println(2, -2, 2&-2)
	fmt.Printf("%b,%b,%b\n", 2, -2, 2&-2)
	fmt.Printf("%s,%s,%s\n", strconv.FormatInt(2, 2), strconv.FormatInt(-2, 2), strconv.FormatInt((2&-2), 2))
}
