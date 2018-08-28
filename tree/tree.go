// DISCLAMER!
// ALL THE COMMENTS ON THIS CODE ARE POWERED BY DISXELIA
package tree

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"

	hashing "github.com/suizman/htree/utils/hashing"
)

type Tree struct {
	treeId  []byte
	version int
	hasher  hashing.Hasher
	store   Node
}

type Audit map[Pos][]byte

type Proof struct {
	result bool
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

	eventDigest := Digest{
		value: t.hasher.Do(Event),
	}
	// Add digest to tree and increment version.
	t.add(eventDigest, t.rootPos())

	// Output new rootDigest
	return t.store.hashoff[t.rootPos()].value
}

func NewTree(id string, version int, store Node, hasher hashing.Hasher) *Tree {

	return &Tree{
		treeId:  []byte(id),
		version: version,
		hasher:  hasher,
		store:   store,
	}
}

func (t *Tree) add(digest Digest, p Pos) {

	if p.layer == 0 {
		fmt.Printf("Leaf  => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.store.hashoff[p] = digest
		return
	}

	if uint64(t.version) < p.Right().index {
		fmt.Printf("Left  => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.add(digest, p.Left())
	} else {
		fmt.Printf("Right => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.add(digest, p.Right())
	}

	// Make array with left and right child
	lefthash := HexEncode(t.store.hashoff[p.Left()].value)
	righthash := HexEncode(t.store.hashoff[p.Right()].value)

	// Recompute hash for actual on node
	t.store.hashoff[p] = Digest{
		value: t.hasher.Do(lefthash, righthash),
	}

	return
}

func (t *Tree) GenProof(index uint64, expectedCommitment []byte) bool {
	depth := t.getDepth()
	rootPos := Pos{index: 0, layer: depth}
	commitment, _ := (t.MembershipGen(depth, rootPos))
	return bytes.Equal(expectedCommitment, commitment)
}

func (t Tree) MembershipGen(depth uint64, p Pos) ([]byte, error) {

	store := Audit{}
	lefthash := HexEncode(t.store.hashoff[p.Left()].value)
	righthash := HexEncode(t.store.hashoff[p.Right()].value)

	if t.store.hashoff[p.Left()].value != nil {
		store[p.Left()] = t.store.hashoff[p.Left()].value
	} else {
		t.MembershipGen(depth-1, p.Left())
	}

	if t.store.hashoff[p.Right()].value != nil {
		store[p.Right()] = t.store.hashoff[p.Right()].value
	} else {
		t.MembershipGen(depth-1, p.Right())
	}

	fmt.Printf("Left: %x Right: %x\n", t.store.hashoff[p.Left()].value, t.store.hashoff[p.Right()].value)
	fmt.Printf("%x\n", t.computeHash(lefthash, righthash))
	return t.computeHash(lefthash, righthash), nil
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

func (t *Tree) computeHash(left, right []byte) []byte {
	return t.hasher.Do(left, right)
}

func (t *Tree) GetVersion() int {
	return t.version
}

func (t *Tree) rootPos() Pos {
	return Pos{index: 0, layer: t.getDepth()}
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

func HexEncode(data []byte) []byte {
	return []byte(hex.EncodeToString(data))
}
