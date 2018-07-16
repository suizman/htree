package tree

import (
	"encoding/binary"
	"fmt"
	"math"

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

func Travel(i uint64, p Pos) bool {

	foreverAlone := Pos{
		index: 0,
		layer: 0,
	}

	if p == foreverAlone {
		fmt.Printf("First node: %v\n", p)
		return true
	}

	if i <= p.layer {
		fmt.Printf("Go left. Index: %v|Layer: %v\n", p.index, p.layer)
		Travel(i, Pos{
			index: p.index,
			layer: p.layer - 1,
		})
	} else {
		fmt.Printf("Go Right. Index: %v|Layer: %v\n", p.index, p.layer)
		Travel(i, Pos{
			index: p.index + pow(2, p.layer) - 1,
			layer: p.layer - 1,
		})
	}
	return true
}

func getDepth(index uint64) uint64 {
	return uint64(math.Ceil(math.Log2(float64(index + 1))))
}

// Utility to calculate arbitraty pow and return an int64
func pow(x, y uint64) uint64 {
	return uint64(math.Pow(float64(x), float64(y)))
}

// uInt64AsBytes returns the []byte representation of a unit64
func uInt64AsBytes(i uint64) []byte {
	valuebytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valuebytes, i)
	return valuebytes
}
