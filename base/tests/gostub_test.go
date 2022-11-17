package tests

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/prashantv/gostub"
)

var now = time.Now

func Now() time.Time {
	return time.Now()
}

func TestStub(t *testing.T) {
	stubs := gostub.Stub(&now, func() time.Time {
		return time.Now().AddDate(1, 0, 1)
	})

	nt, st := time.Now(), now()
	fmt.Println("normal:", nt)
	fmt.Println("stub  :", st)

	if m := nt.Sub(st).Hours() / 24 / 30; math.Abs(m) < 12 {
		t.Error("stub error")
	} else {
		fmt.Println("month:", m)
	}
	defer stubs.Reset()
}
