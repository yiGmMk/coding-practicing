package base

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"golang.org/x/crypto/sha3"
)

func TestSHA(t *testing.T) {
	data := []byte("a gopher")
	data1 := []byte("a gopher ")
	t.Run("SHA-256", func(t *testing.T) {
		sum := sha256.Sum256(data)
		sum1 := sha256.Sum256(data1)
		fmt.Printf("sha-2-256:%x\nsha-2-256:%x\n", sum, sum1)
	})

	t.Run("SHA-3", func(t *testing.T) {
		sum := sha3.Sum256(data)
		sum1 := sha3.Sum256(data1)
		fmt.Printf("sha-3-256:%x\nsha-3-256:%x\n", sum, sum1)
	})
}
