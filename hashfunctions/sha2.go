package hashfunctions

import (
	"crypto/sha256"
	"crypto/sha512"
)

func SHA2_256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func SHA2_512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}
