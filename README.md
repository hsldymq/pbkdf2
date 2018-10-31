# 闲得蛋疼自己写了个PBKDF2的实现

根据[RFC 8018](https://tools.ietf.org/html/rfc8018)中5.2的描述,再参(chao)考(xi)[官方实现](https://github.com/golang/crypto/blob/master/pbkdf2/pbkdf2.go),自己再这个实现上小动了两刀,性能无差异,只是为了巩固一下印象而已.

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