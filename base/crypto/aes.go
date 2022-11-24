package base

import (
	"crypto/aes"
	"crypto/cipher"
)

const (
	// aes 密钥(key) 32字节 => AES-256
	_aesKey = "12345678123456781234567812345678"
	// 这里初始变量采用固定字符串，这样多次运行结果相同，便于示例说明
	_aesInitVariables = "abcdefghijklmnop" //iv
)

func AesEncrypt(data string) ([]byte, error) {
	key := []byte(_aesKey)

	// 创建AES分组密码算法实例
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 待加密的明文字符串（长度恰为分组长度的整数倍）
	plaintext := []byte(data)

	// 存储密文的切片，预留出在密文头部放置初始变量的空间
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// 这里初始变量采用固定字符串，这样多次运行结果相同，便于示例说明
	iv := []byte(_aesInitVariables)

	// 创建分组模式的实例，这里使用CBC模式
	cbcModeEncrypter := cipher.NewCBCEncrypter(aesCipher, iv)

	// 对明文进行加密
	cbcModeEncrypter.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 这里将初始变量放在密文的头部（初始变量的长度 = block length）
	copy(ciphertext[:aes.BlockSize], []byte(_aesInitVariables))

	return ciphertext, nil
}

func AesDecrypt(data string) ([]byte, error) {
	ciphertextWithIV := []byte(data)
	// 从密文数据中提取初始向量数据
	iv := ciphertextWithIV[:aes.BlockSize]

	// 待解密的密文数据
	ciphertext := ciphertextWithIV[aes.BlockSize:]

	key := []byte(_aesKey)
	// 创建AES分组密码算法实例
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 解密后存放明文的字节切片
	plaintext := make([]byte, len(ciphertext))

	// 创建分组模式的实例，这里使用CBC模式
	cbcModeDecrypter := cipher.NewCBCDecrypter(aesCipher, iv)

	// 对密文进行解密
	cbcModeDecrypter.CryptBlocks(plaintext, ciphertext)
	// 支持原地替换,直接利用ciphertext存储解密后的明文 cbcModeDecrypter.CryptBlocks(ciphertext, ciphertext)
	return plaintext, nil
}
