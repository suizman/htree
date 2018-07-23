package hasher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasher(t *testing.T) {

	expectedDigest := []byte{0xc7, 0xbe, 0x1e, 0xd9, 0x2, 0xfb, 0x8d, 0xd4, 0xd4, 0x89, 0x97, 0xc6, 0x45, 0x2f, 0x5d, 0x7e, 0x50, 0x9f, 0xbc, 0xdb, 0xe2, 0x80, 0x8b, 0x16, 0xbc, 0xf4, 0xed, 0xce, 0x4c, 0x7, 0xd1, 0x4e}

	hasher := new(Sha256Hasher)
	digest := hasher.Do([]byte("This is a test"))
	require.Equal(t, expectedDigest, digest, "Invalid digest:")
	t.Logf("\nExpected digest: %x\nGenerated digest:%x\n", expectedDigest, digest)
}

func TestNewHashahin(t *testing.T) {

	expectedDigest := []byte{0xbd, 0x81, 0x43, 0xdb, 0x4e, 0x6c, 0x9e, 0xa6, 0xe3, 0x88, 0x19, 0x31, 0xc3, 0x5b, 0x71, 0xda, 0xd7, 0xe0, 0x74, 0xff, 0xbe, 0x5d, 0x6e, 0xc, 0xa8, 0xe7, 0xfe, 0x18, 0x86, 0xf4, 0x7b, 0x60}

	hasher := NewHashahin()
	hasher.h.Write([]byte("Hello Mr. Hashashin."))
	digest := hasher.h.Sum(nil)

	require.Equal(t, expectedDigest, digest, "Invalid digest:")
	t.Logf("\nExpected digest: %x\nGenerated digest:%x\n", expectedDigest, digest)
}

func BenchmarkHasher(b *testing.B) {

	hasher := new(Sha256Hasher)

	for i := 0; i < b.N; i++ {
		event := fmt.Sprintf("Test event: %v", i)
		hasher.Do([]byte(event))
	}
}

func BenchmarkQedHasher(b *testing.B) {

	hasher := new(Sha256Hasher)
	b.N = 50000000

	for i := 0; i < b.N; i++ {
		event := fmt.Sprintf("Test event: %v", i)
		hasher.Do([]byte(event))
	}
}

func BenchmarkQedHasherV2(b *testing.B) {

	hasher := NewHashahin()
	b.N = 50000000

	for i := 0; i < b.N; i++ {
		hasher.h.Write([]byte("Hello Mr. Hashashin."))
		hasher.h.Sum(nil)
	}
}
