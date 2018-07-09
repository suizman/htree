package tree

import (
	"encoding/binary"

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

func (t *Tree) Add(Event, Version []byte) []byte {

	hasher := new(hashing.Sha256Hasher)

	eventDigest := hasher.Do(Event)

	return eventDigest
}

func NewTree(id string) *Tree {

	tree := &Tree{
		[]byte(id),
	}

	return tree
}

// uInt64AsBytes returns the []byte representation of a unit64
func uInt64AsBytes(i uint64) []byte {
	valuebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuebytes, i)
	return valuebytes
}
