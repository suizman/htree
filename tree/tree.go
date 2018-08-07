// DISCLAMER!
// ALL THE COMMENTS ON THIS CODE ARE POWERED BY DISXELIA
package tree

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
)

type Tree struct {
	treeId  []byte
	version uint64
	hasher  hash.Hash
	store   Node
}

type Proof struct {
	result bool
}

type Audit struct {
	path map[string][]byte
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

	t.hasher.Write(Event)

	rootDigest := Digest{
		value: t.hasher.Sum(nil),
	}

	rootPos := Pos{
		index: 0,
		layer: t.getDepth(),
	}

	// Add root digest to tree and increment version.
	t.add(rootDigest, rootPos)
	t.version++
	return rootDigest.value
}

func NewTree(id string, version uint64, store Node) *Tree {

	return &Tree{
		treeId:  []byte(id),
		version: version,
		hasher:  sha256.New(),
		store:   store,
	}

}

func (t *Tree) add(digest Digest, p Pos) {

	if p.layer == 0 {
		fmt.Printf("Leaf node  => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.store.hashoff[p] = digest
		return
	}

	if t.version <= p.index+pow(2, p.layer-1) {
		fmt.Printf("Go left    => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.add(digest, p.Left())
	} else {
		fmt.Printf("Go right   => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.add(digest, p.Right())
	}

	// Make array with left and right child
	rl := make([]byte, 2*sha256.Size)
	copy(rl, t.store.hashoff[p.Left()].value)
	rl = append(rl, t.store.hashoff[p.Right()].value...)

	// Recompute hash for actual on node
	t.hasher.Write(rl)
	t.store.hashoff[p] = Digest{
		value: []byte(t.hasher.Sum(nil)),
	}
	return
}

func (p *Proof) GenProof() bool {
	return p.result
}

func (t Tree) AuditPath(index, version uint64) *Tree {

	return &Tree{}
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

func (t *Tree) GetVersion() uint64 {
	return t.version
}

// v = tree version
func (t *Tree) Travel(p Pos) {

	if p.layer == 0 {
		fmt.Printf("Leaf node  => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		return
	}

	if t.version <= p.index+pow(2, p.layer-1) {
		fmt.Printf("Go left    => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.Travel(p.Left())
	} else {
		fmt.Printf("Go right   => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.Travel(p.Right())
	}

	return
}

func (t *Tree) getDepth() uint64 {
	return uint64(math.Ceil(math.Log2(float64(t.version + 1))))
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
