package base

import (
	"crypto/hmac"
	"crypto/sha256"
)

var (
	_hMacKey = "12345678123456781234567812345678" //密钥 32字节
)

func HMac(data []byte) ([]byte, error) {
	// 创建hmac实例（使用SHA-256单向散列函数）
	mac := hmac.New(sha256.New, []byte(_hMacKey))
	_, err := mac.Write(data)
	if err != nil {
		return nil, err
	}

	// 计算mac值
	m := mac.Sum(nil)
	return m, nil
}
