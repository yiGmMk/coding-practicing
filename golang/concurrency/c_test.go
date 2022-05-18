package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestC(t *testing.T) {
	j := 0
	for {
		fmt.Println(time.Now())
		j++
		if j > 1000000 {
			break
		}
	}
	Range()
	fmt.Println("end done")
}
