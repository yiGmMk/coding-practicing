package base

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/pkg/errors"
)

// rsa 填充方案:
// 0.默认填充方案:PKCS#1 v1.5
// 1.RSA-OAEP(Optimal Asymmetric Encryption Padding，最优非对称加密填充),被认为是一种可信赖、满足强度要求的填充方案

// rsa秘钥
var _defaultRsaKey, err = rsa.GenerateKey(rand.Reader, 2028)

// 加密
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
	// rsa.EncryptOAEP和rsa.DecryptOAEP的第二个参数都是一个随机数生成器（这里传入rand.Reader），
	// RSA-OAEP会通过随机数使每次生成的密文呈现不同的排列方式，因此多次运行上述示例程序所得到的密文结果都是不同的
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
