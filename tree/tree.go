package tree

import (
	"encoding/binary"
	"fmt"

	hashing "github.com/suizman/htree/utils/hashing"
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

func (t *Tree) Add(Event, Version []byte) ([]byte, error) {

	hasher := new(hashing.Sha256Hasher)

	eventDigest := hasher.Do(Event)
	rootDigest, err := eventDigest, Version

	if err != nil {
		return nil, fmt.Errorf("Unable to add event %v", err)
	}

	return rootDigest, nil
}

// uInt64AsBytes returns the []byte representation of a unit64
func uInt64AsBytes(i uint64) []byte {
	valuebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuebytes, i)
	return valuebytes
}
