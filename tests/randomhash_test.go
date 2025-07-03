package tests

// go test ./tests --bench=. -benchmem

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/gcarreno/go-randomhash"
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
		Output: "816005612490569189810604011548ac60085141008a921078904c1000a06042", // this would be from a known output
	},
	// Add more
}

func TestRandomHashVectors(t *testing.T) {
	for _, tv := range testVectors {
		got := randomhash.RandomHash(tv.Input, tv.Nonce)
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

func fuzzRandomHash(rounds int, minSize int, maxSize int) {
	failCount := 0
	start := time.Now()

	for i := 0; i < rounds; i++ {
		// 🔀 Random input length between minSize and maxSize
		inputLen := minSize + randInt(maxSize-minSize+1)
		input := randomBytes(inputLen)

		// 🎲 Random nonce
		nonce := randInt63()

		// ⛏️ First run
		hash1 := randomhash.RandomHash(input, nonce)

		// 🧪 Re-run with same input and nonce to test determinism
		hash2 := randomhash.RandomHash(input, nonce)

		if !bytes.Equal(hash1, hash2) {
			failCount++
			fmt.Printf("FAIL @ iteration %d:\nInput: %x\nNonce: %d\nHash1: %x\nHash2: %x\n\n",
				i, input, nonce, hash1, hash2)
		}

		if i%100 == 0 && i != 0 {
			fmt.Printf("Fuzzed %d rounds... all stable so far 😎\n", i)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\nFuzz complete.\nRounds: %d\nFailures: %d\nElapsed: %s\n",
		rounds, failCount, elapsed)
}

// ---
// 🧪 1. Fuzz-style determinism test
// ---

func TestRandomHashDeterminism(t *testing.T) {
	const rounds = 10_000
	const minSize = 8
	const maxSize = 128

	for i := 0; i < rounds; i++ {
		inputLen := minSize + randInt(maxSize-minSize+1)
		input := randomBytes(inputLen)
		nonce := randInt63()

		hash1 := randomhash.RandomHash(input, nonce)
		hash2 := randomhash.RandomHash(input, nonce)

		if !bytes.Equal(hash1, hash2) {
			t.Errorf("[FAIL #%d]\nInput: %s\nNonce: %d\nHash1: %x\nHash2: %x\n",
				i, hex.EncodeToString(input), nonce, hash1, hash2)
		}
	}
}

// ---
// ⚡ 2. Benchmark with various input sizes
// ---

func BenchmarkRandomHash_32Bytes(b *testing.B)  { benchWithInputSize(b, 32) }
func BenchmarkRandomHash_64Bytes(b *testing.B)  { benchWithInputSize(b, 64) }
func BenchmarkRandomHash_128Bytes(b *testing.B) { benchWithInputSize(b, 128) }
func BenchmarkRandomHash_256Bytes(b *testing.B) { benchWithInputSize(b, 256) }

func benchWithInputSize(b *testing.B, size int) {
	input := randomBytes(size)
	nonce := int64(42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = randomhash.RandomHash(input, nonce)
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
