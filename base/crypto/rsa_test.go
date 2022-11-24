package base

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err == nil {
		fmt.Println(key.PublicKey)
	}
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
