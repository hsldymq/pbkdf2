package pbkdf2

import (
	"crypto/hmac"
	"encoding/binary"
	"hash"
)

type pbkdf2 struct {
	hashMaker func() hash.Hash
}

func New(h func() hash.Hash) *pbkdf2 {
	return &pbkdf2{
		hashMaker: h,
	}
}

func (h *pbkdf2) MakeKey(password []byte, salt []byte, iterCount int, dkLen int) []byte {
	prf := hmac.New(h.hashMaker, password)
	hLen := prf.Size()
	// l = CEIL(dkLen / hLen)
	l := (dkLen + hLen - 1) / hLen
	key := make([]byte, 0, l*hLen)
	U := make([]byte, hLen)
	// INT (i)
	inti := make([]byte, 4)

	for i := 1; i <= l; i++ {
		// U_1 = PRF (P, S || INT (i))
		binary.BigEndian.PutUint32(inti, uint32(i))
		prf.Reset()
		prf.Write(salt)
		prf.Write(inti)
		key = prf.Sum(key)
		T := key[(i-1)*hLen:]
		copy(U, T)

		// U_2 = PRF (P, U_1) ,
		// ...
		// U_c = PRF (P, U_{c-1})
		for c := 1; c < iterCount; c++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for idx := range U {
				// Ti = U_1 \xor U_2 \xor ... \xor U_c
				T[idx] ^= U[idx]
			}
		}
	}

	return key[:dkLen]
}
