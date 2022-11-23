package base

import (
	"fmt"
	"testing"
)

func TestHMac(t *testing.T) {
	fmt.Printf("key.len=%d,%d\n", len(_hMacKey), len("12345678123456781234567812345678"))

	msg := "favorite programming language:go"
	mac, err := HMac([]byte(msg))
	if err != nil {
		t.Error(err)
	}

	ms := fmt.Sprintf("%x", mac) // mac到string
	fmt.Printf("mac值 = %s\nmac值 = %s", ms, mac)
}
