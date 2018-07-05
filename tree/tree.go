package tree

import (
	"encoding/binary"
	"fmt"
)

type Tree struct {
	treeId []byte
}

type Node struct {
	Leaf uint64
}

type Event struct {
	Event []byte
}

func (t *Tree) Add(Digest, Version []byte) ([]byte, error) {

	rootDigest, err := Digest, Version
	if err != nil {
		fmt.Errorf("Unable to add event %v:\n", err)
	}

	return rootDigest, nil
}

// uInt64AsBytes returns the []byte representation of a unit64
func uInt64AsBytes(i uint64) []byte {
	valuebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuebytes, i)
	return valuebytes
}
