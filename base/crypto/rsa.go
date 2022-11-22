package base

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/pkg/errors"
)

var _defaultRsaKey, err = rsa.GenerateKey(rand.Reader, 2028)

// 加密
// 填充方案: RSA-OAEP(Optimal Asymmetric Encryption Padding，最优非对称加密填充),被认为是一种可信赖、满足强度要求的填充方案
func RsaEncrypt(data []byte, publicKey ...*rsa.PublicKey) ([]byte, error) {
	var key *rsa.PublicKey
	if len(publicKey) > 0 {
		key = publicKey[0]
	} else {
		key = &_defaultRsaKey.PublicKey
	}
	if key == nil {
		return nil, errors.Errorf("no key provided")
	}
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, key, data, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func RsaDecrypt(data []byte, privateKey ...*rsa.PrivateKey) ([]byte, error) {
	var key *rsa.PrivateKey
	if len(privateKey) > 0 {
		key = privateKey[0]
	} else {
		key = _defaultRsaKey
	}
	if key == nil {
		return nil, errors.Errorf("no key provided")
	}
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, data, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
