package tree

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/suizman/htree/utils/hashing"
)

type Tree struct {
	treeId  []byte
	version uint64
	hasher  hasher.Hasher
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

func (t *Tree) Add(Event []byte) []byte {

	t.version++
	hasher := new(hasher.Sha256Hasher)

	eventDigest := hasher.Do(Event)

	return eventDigest
}

func NewTree(id string, version uint64, hasher hasher.Hasher, store Node) *Tree {

	return &Tree{
		[]byte(id),
		version,
		hasher,
		store,
	}

}

func (t *Tree) add(digest Digest, p Pos) {

	if p.layer == 0 {
		fmt.Printf("Leaf node: %v\n", p)
		t.store.hashoff[p] = digest
		return
	}

	if t.version <= p.index+pow(2, p.layer-1) {
		fmt.Printf("Go left. Index: %v|Layer: %v\n", p.index, p.layer)
		t.add(digest, p.Left())
	} else {
		fmt.Printf("Go Right. Index: %v|Layer: %v\n", p.index, p.layer)
		t.add(digest, p.Right())
	}

	l, r := []byte(fmt.Sprint(t.store.hashoff[p.Left()])), []byte(fmt.Sprint(t.store.hashoff[p.Right()]))
	t.hasher.Do(l, r)
	// t.hasher.Do(t.store.hashoff[p.Left()], t.store.hashoff[p.Right()])
	return
}

func (p *Pos) Left() Pos {

	return Pos{
		index: p.index,
		layer: p.layer - 1,
	}

}

func (p *Pos) Right() Pos {

	return Pos{
		index: p.index + pow(2, p.layer-1),
		layer: p.layer - 1,
	}

}

// v = tree version
func (t *Tree) Travel(p Pos) {

	if p.layer == 0 {
		fmt.Printf("Leaf node: %v\n", p)
		return
	}

	if t.version <= p.index+pow(2, p.layer-1) {
		fmt.Printf("Go left. Index: %v|Layer: %v\n", p.index, p.layer)
		t.Travel(p.Left())
	} else {
		fmt.Printf("Go Right. Index: %v|Layer: %v\n", p.index, p.layer)
		t.Travel(p.Right())
	}

	return
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
