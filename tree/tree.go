package tree

import (
	"encoding/binary"

	hashing "github.com/suizman/htree/utils/hashing"
)

type Tree struct {
	treeId  []byte
	version uint64
	node    Node
}
type Node struct {
	position map[Pos]Digest
}
type Digest struct {
	value []byte
}

type Event struct {
	event []byte
}

type Pos struct {
	index uint64
	layer uint64
}

func (t *Tree) Add(Event, version []byte) []byte {

	hasher := new(hashing.Sha256Hasher)

	eventDigest := hasher.Do(Event)

	return eventDigest
}

func NewTree(id string, version uint64, node Node) *Tree {

	tree := &Tree{
		[]byte(id),
		version,
		node,
	}

	return tree
}

// uInt64AsBytes returns the []byte representation of a unit64
func uInt64AsBytes(i uint64) []byte {
	valuebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuebytes, i)
	return valuebytes
}
