package base

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/pkg/errors"
)

var _rsaPssKey, _ = rsa.GenerateKey(rand.Reader, 2048)

func RsaPssSign(data []byte) ([]byte, []byte, error) {
	if _rsaPssKey == nil {
		return nil, nil, errors.Errorf("no key provided")
	}

	// 计算摘要
	digest := sha256.Sum256(data)

	// 私钥签名
	sign, err := rsa.SignPSS(rand.Reader, _rsaPssKey, crypto.SHA256, digest[:], nil)
	if err != nil {
		return nil, nil, err
	}
	return sign, digest[:], nil
}

func RsaPassVerify(sign, digest []byte) error {
	return rsa.VerifyPSS(&_rsaPssKey.PublicKey, crypto.SHA256, digest, sign, nil)
}
