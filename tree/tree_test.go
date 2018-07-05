package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {

	tree := &Tree{[]byte("tree-fake-id")}

	input := []struct {
		eventDigest    []byte
		expectedDigest []byte
	}{
		{[]byte{0x00}, []byte{0x00}},
	}

	for v, c := range input {
		version := uint64(v)
		rh, err := tree.Add(c.eventDigest, uInt64AsBytes(version))
		require.NoError(t, err, "Unable to add event to tree")
		require.Equalf(t, c.expectedDigest, rh, "Incorrect root hash for index %d", v)
	}
}
