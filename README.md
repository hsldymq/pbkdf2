# 闲得蛋疼自己写了个PBKDF2的实现

PBKDF2是一种密钥导出算法,通过对输入的密码与盐进行若干次迭代运算,增加了导出密钥的时间,以及其破解的难度.

PBKDF2定义于RFC 8018中,它的前任是RFC 2898.两个文档对于PBKDF2部分的定义和描述并无区别.

这个实现根据[RFC 8018](https://tools.ietf.org/html/rfc8018)中5.2的描述,再参(chao)考(xi)[官方实现](https://github.com/golang/crypto/blob/master/pbkdf2/pbkdf2.go),自己再这个实现上小动了两刀,性能无差异,只是为了巩固一下印象而已.

### Example
```golang

import (
    "github.com/hsldymq/pbkdf2"
    "crypto/sha256"
)

// ...

df := pbkdf2.New(sha256.New)
key := (df.MakeKey([]byte(password), []byte(salt), iter, dkLen))

// do something

```