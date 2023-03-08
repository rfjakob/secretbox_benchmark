package secretbox_benchmark

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/nacl/secretbox"
)

func benchmarkSecretboxSealSize(b *testing.B, size int) {
	message := make([]byte, size)
	out := make([]byte, size+secretbox.Overhead)
	var nonce [24]byte
	var key [32]byte

	b.SetBytes(int64(size))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out = secretbox.Seal(out[:0], message, &nonce, &key)
	}
}

func BenchmarkSecretbox(b *testing.B) {
	for blockSize := 1024; blockSize <= 1048576; blockSize *= 2 {
		name := fmt.Sprintf("blockSize%d", blockSize)
		b.Run(name, func(b *testing.B) { benchmarkSecretboxSealSize(b, blockSize) })
	}
}
