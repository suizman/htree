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

func BenchmarkHasher(b *testing.B) {

	hasher := new(Sha256Hasher)

	for i := 0; i < b.N; i++ {
		event := fmt.Sprintf("Test event: %v", i)
		hasher.Do([]byte(event))
	}
}
