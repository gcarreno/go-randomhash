package hashfunctions

import "crypto/sha3"

func SHA3_256(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}
