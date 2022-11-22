# Crypto

## 分类

### 分组密码（block cipher）

主流应用算法,仅支持处理固定长度数据块,包含:DES,AES等

分组密码算法每次仅加密一个明文分组。如果明文因总长度超过分组长度而存在多个分组，那么分组密码算法会被迭代调用以逐个处理明文分组。迭代的方法称为分组密码算法的模式

### 流密码（stream cipher）

流密码是对数据流进行连续处理的一类算法

## 非对称加密/公钥密码

私有密钥（private key，简称私钥）和公共密钥（public key，简称公钥）.
RSA是世界上使用最广泛的公钥密码算法，Go语言的crypto/rsa包提供了对RSA算法的实现

```go
// $GOROOT/src/crypto/rsa/rsa.go
// RSA的公钥可以看成数对(E, N)，而私钥可以看出数对(D, N)
// 密文= 明文^E mod N (RSA密文就是对代表明文的数字的E次方求mod N的结果)
// 明文= 密文^D mod N
type PublicKey struct {
    N *big.Int // modulus 模数
    E int      // public exponent 指数
}

type PrivateKey struct {
    PublicKey            // 公钥部分
    D         *big.Int   // 私钥组件
    Primes    []*big.Int // N的质因数，有2个或2个以上元素
    ...
}
```
