package randomhash

import (
	"math/rand"

	"github.com/gcarreno/go-randomhash/hashfunctions"
)

type (
	HashVersion int
	HashFunc    func([]byte) []byte
)

const (
	Version1 HashVersion = iota
	// Version2

	DefaultHashRound = 5
)

var (
	hashFuncs = []HashFunc{
		hashfunctions.SHA2_256,
		hashfunctions.SHA2_512,
		hashfunctions.BLAKE2B_512,
		hashfunctions.SHA3_256,
	}
)

func xorBytes(a, b []byte) []byte {
	n := min(len(a), len(b))
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func andBytes(a, b []byte) []byte {
	n := min(len(a), len(b))
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = a[i] & b[i]
	}
	return result
}

func rotateLeft(data []byte, n uint) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[i] = (b << n) | (b >> (8 - n))
	}
	return result
}

func reverseBytes(data []byte) []byte {
	n := len(data)
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = data[n-1-i]
	}
	return result
}

func flipBits(data []byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[i] = ^b
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shuffledHashFuncs(seed int64) []HashFunc {
	r := rand.New(rand.NewSource(seed))
	clone := make([]HashFunc, len(hashFuncs))
	copy(clone, hashFuncs)
	r.Shuffle(len(clone), func(i, j int) {
		clone[i], clone[j] = clone[j], clone[i]
	})
	return clone
}

func RandomHash(input []byte, nonce int64, version HashVersion) []byte {
	switch version {
	case Version1:
		return randomHashVersion1(input, nonce)
	default:
		return nil
	}
}

func randomHashVersion1(input []byte, nonce int64) []byte {
	rnd := rand.New(rand.NewSource(nonce))
	funcs := shuffledHashFuncs(nonce)
	state := input
	prev := input
	steps := DefaultHashRound + rnd.Intn(DefaultHashRound)

	for i := 0; i < steps; i++ {
		hfn := funcs[i%len(funcs)]
		hashed := hfn(state)

		// Mix it up with one of FIVE juicy entropy operations
		switch rnd.Intn(5) {
		case 0:
			hashed = xorBytes(hashed, prev)
		case 1:
			hashed = andBytes(hashed, prev)
		case 2:
			rot := uint(rnd.Intn(7) + 1)
			hashed = rotateLeft(hashed, rot)
		case 3:
			hashed = reverseBytes(hashed)
		case 4:
			hashed = flipBits(hashed)
		}

		prev = state
		state = hashed
	}

	return state
}
