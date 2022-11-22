package base

import (
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	text := "a gopher"
	encryptData, err := RsaEncrypt([]byte(text))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("text   :%s\nencrypt:%x\n", text, encryptData)

	decrypt, err := RsaDecrypt(encryptData)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decrypt:%s\n", decrypt)
}
