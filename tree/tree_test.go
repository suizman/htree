package tree

import (
	"fmt"
	"testing"
)

// func TestAdd(t *testing.T) {

// 	tree := &Tree{[]byte("tree-fake-id")}

// 	input := []struct {
// 		eventDigest    []byte
// 		expectedDigest []byte
// 	}{
// 		{[]byte{0x00}, []byte{0x00}},
// 	}

// 	for v, c := range input {
// 		version := uint64(v)
// 		rh, err := tree.Add(c.eventDigest, uInt64AsBytes(version))
// 		require.NoError(t, err, "Unable to add event to tree")
// 		require.Equalf(t, c.expectedDigest, rh, "Incorrect root hash for index %d", v)
// 	}
// }

func TestNewAdd(t *testing.T) {

	tree := &Tree{[]byte("tree-fake-id")}
	event := []byte("Test event")
	// expectedDigest := []byte{0x1}
	_, digest := tree.Add(event, []byte("0"))

	// if err != nil {
	// 	t.Errorf("Oops %v", err)
	// }
	fmt.Printf("%x", digest)
}
