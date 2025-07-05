package tests

// go test ./tests --bench=. -benchmem

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"

	randomhash "github.com/gcarreno/go-randomhash"
)

type TestVector struct {
	Input  []byte
	Nonce  int64
	Output string // hex-encoded expected hash
}

var testVectors = []TestVector{
	{
		Input:  []byte("hello"),
		Nonce:  12345,
		Output: "bc5760421118464abc929b3811555a8669cacdcf41378500c5e1a848628336f0", // this would be from a known output
	},
	{
		Input:  []byte("Hello, World!"),
		Nonce:  12345,
		Output: "18bf19d7f24c8689f5eb71b9710225a05681b05304698c3c8ef1e4ab50f68025", // this would be from a known output
	},
	{
		Input:  []byte("Hello, World! 😎"),
		Nonce:  12345,
		Output: "7922ea47a65a950f18b0e3827a7e8de091e1d8585644fbe7e8268865631adb64", // this would be from a known output
	},
	// Add more
}

func TestRandomHashVectorsWithVersion1(t *testing.T) {
	for _, tv := range testVectors {
		got := randomhash.RandomHash(tv.Input, tv.Nonce, randomhash.Version1)
		if hex.EncodeToString(got) != tv.Output {
			t.Errorf("Input: %q Nonce: %d\nExpected: %s\nGot:      %x",
				tv.Input,
				tv.Nonce,
				tv.Output,
				got,
			)
		}
	}
}

// ---
// 🧪 1. Fuzz-style determinism test
// ---

func TestRandomHashDeterminismWithVersion1(t *testing.T) {
	const rounds = 10_000
	const minSize = 8
	const maxSize = 128

	for i := 0; i < rounds; i++ {
		inputLen := minSize + randInt(maxSize-minSize+1)
		input := randomBytes(inputLen)
		nonce := randInt63()

		hash1 := randomhash.RandomHash(input, nonce, randomhash.Version1)
		hash2 := randomhash.RandomHash(input, nonce, randomhash.Version1)

		if !bytes.Equal(hash1, hash2) {
			t.Errorf("[FAIL #%d]\nInput: %s\nNonce: %d\nHash1: %x\nHash2: %x\n",
				i, hex.EncodeToString(input), nonce, hash1, hash2)
		}
	}
}

// ---
// ⚡ 2. Benchmark with various input sizes
// ---

func BenchmarkRandomHash_32Bytes(b *testing.B)  { benchWithInputSizeWithVersion1(b, 32) }
func BenchmarkRandomHash_64Bytes(b *testing.B)  { benchWithInputSizeWithVersion1(b, 64) }
func BenchmarkRandomHash_128Bytes(b *testing.B) { benchWithInputSizeWithVersion1(b, 128) }
func BenchmarkRandomHash_256Bytes(b *testing.B) { benchWithInputSizeWithVersion1(b, 256) }

func benchWithInputSizeWithVersion1(b *testing.B, size int) {
	input := randomBytes(size)
	nonce := int64(42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = randomhash.RandomHash(input, nonce, randomhash.Version1)
	}
}

// ---
// 🔧 3. Random utilities
// ---

func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func randInt(n int) int {
	b := make([]byte, 2)
	rand.Read(b)
	return int(b[0]) % n
}

func randInt63() int64 {
	var b [8]byte
	rand.Read(b[:])
	return int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 |
		int64(b[3])<<32 | int64(b[4])<<24 | int64(b[5])<<16 |
		int64(b[6])<<8 | int64(b[7])
}
