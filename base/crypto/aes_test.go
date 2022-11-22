package base

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	// 待加密的明文字符串（长度恰为分组长度的整数倍）
	plaintext := "I love go programming language!!"
	// 存储密文的切片，预留出在密文头部放置初始变量的空间
	ciphertext, err := AesEncrypt(plaintext)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("明文：%s\n", plaintext)
	fmt.Printf("密文(包含IV)：%x\n", ciphertext)

	// 用%x打印的字符串与%s不一样
	// 带有初始向量的密文数据（前16字节为初始向量）
	data := "6162636465666768696a6b6c6d6e6f70bc93b5cb1a081b47357f73d40966e3ce53c29db21a13bec2f9be4f76d8f09f2b" //%v
	ciphertextWithIV, err := hex.DecodeString(string(data))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("密文,decode from hex: %x\n", ciphertextWithIV)

	//-------------解密-------------
	text, err := AesDecrypt(string(ciphertext))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("密文(包含IV)：%x\n", ciphertext)
	fmt.Printf("明文：%s\n", text)
}
