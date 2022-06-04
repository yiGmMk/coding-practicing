package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

// system status
// cpu/memory/disk/network etc.
func TestGetOSSatt(t *testing.T) {
	before, err := cpu.Get()
	if err != nil {
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return
	}
	total := float64(after.Total - before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}
