package base

import (
	"fmt"
	"testing"
)

func TestRsaPss(t *testing.T) {
	msg := []byte("2022-11-23,msg,content:a gopher")
	sign, digest, err := RsaPssSign(msg)
	if err != nil {
		t.Fatalf("sign failed,%+v", err)
	}

	fmt.Printf("signature:%s\ndigest:%s\n", fmt.Sprintf("%x", sign), fmt.Sprintf("%x", digest))

	err = RsaPassVerify(sign, digest)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("verify success")
}
