package tests

import (
	"github.com/goal-web/encryption"
	"testing"
)

/**
goos: darwin
goarch: amd64
pkg: github.com/goal-web/encryption/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkAES
BenchmarkAES-4   	 1773345	       633.8 ns/op
*/
func BenchmarkAES(b *testing.B) {
	aes := encryption.AES("123456781234567812345678")
	for i := 0; i < b.N; i++ {
		encrypted := aes.Encode("goal")
		_, _ = aes.Decode(encrypted)

	}
}
