package hashfunctions

import "golang.org/x/crypto/blake2b"

func BLAKE2B_512(data []byte) []byte {
	h, _ := blake2b.New512(nil)
	h.Write(data)
	return h.Sum(nil)
}
