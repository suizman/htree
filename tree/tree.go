package tree

import (
	"encoding/binary"

	"github.com/bbva/qed/hashing"
)

type Tree struct {
	treeId  []byte
	version uint64
	store   Node
}
type Node struct {
	hashoff map[Pos]Digest
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

	eventDigest := hasher.Do(Event, version)

	position := Pos{
		index: 0,
		layer: 0,
	}

	t.store.hashoff[position] = Digest{value: eventDigest}

	return eventDigest
}

func NewTree(id string, version uint64, store Node) *Tree {

	return &Tree{[]byte(id), version, store}

}

// uInt64AsBytes returns the []byte representation of a unit64
func uInt64AsBytes(i uint64) []byte {
	valuebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuebytes, i)
	return valuebytes
}
