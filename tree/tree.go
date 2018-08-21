// DISCLAMER!
// ALL THE COMMENTS ON THIS CODE ARE POWERED BY DISXELIA
package tree

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"math"
)

type Tree struct {
	treeId  []byte
	version int
	hasher  hash.Hash
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
	t.hasher.Write(Event)

	rootDigest := Digest{
		value: t.hasher.Sum(nil),
	}

	// Add root digest to tree and increment version.
	t.add(rootDigest, t.rootPos())
	return rootDigest.value
}

func NewTree(id string, version int, store Node) *Tree {

	return &Tree{
		treeId:  []byte(id),
		version: version,
		hasher:  sha256.New(),
		store:   store,
	}

}

func (t *Tree) add(digest Digest, p Pos) {

	// fmt.Println(t.version)
	if p.layer == 0 {
		fmt.Printf("Leaf node  => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		t.store.hashoff[p] = digest
		return
	}

	// fmt.Println(p.index, p.layer)
	if uint64(t.version) <= p.index+pow(2, p.layer-1) {
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

func (t *Tree) GenProof(index uint64, commitment []byte) bool {
	depth := t.getDepth()
	rootPos := Pos{index: 0, layer: depth}
	expectedCommitment, _ := (t.MembershipGen(depth, rootPos))
	expectedCommitment = []byte(expectedCommitment)
	return bytes.Equal(expectedCommitment, commitment)
}

func (t Tree) MembershipGen(depth uint64, p Pos) ([]byte, error) {
	store := Audit{}
	digest := Digest{value: []byte("digest")}
	if p.index < 0 || p.index > uint64(t.version) {
		return digest.value, errors.New("Invalid index, 0 <= index <= version")
	}

	if p.index == 0 && p.layer == 0 {
		store[Pos{index: 0, layer: 0}] = t.store.hashoff[Pos{index: 0, layer: 0}].value
		return computeHash(digest.value, digest.value), nil
	}

	if t.store.hashoff[p.Left()].value != nil {
		store[p.Left()] = t.store.hashoff[p.Left()].value
		t.MembershipGen(depth-1, p.Left())
	}

	if t.store.hashoff[p.Right()].value != nil {
		store[p.Right()] = t.store.hashoff[p.Right()].value
		t.MembershipGen(depth-1, p.Right())
	}

	return computeHash(digest.value, digest.value), nil
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

func Even(number uint64) bool {
	return number%2 == 0
}

func Odd(number uint64) bool {
	// Odd should return not even.
	// ... We cannot check for 1 remainder.
	// ... That fails for negative numbers.
	return !Even(number)
}

func computeHash(left, right []byte) []byte {
	return []byte(left)
}

func (t *Tree) GetVersion() int {
	return t.version
}

func (t *Tree) rootPos() Pos {
	return Pos{index: 0, layer: t.getDepth()}
}

// v = tree version
func (t *Tree) Travel(p Pos) {

	if p.layer == 0 {
		fmt.Printf("Leaf node  => Index: %v | Layer: %v | Version: %v\n", p.index, p.layer, t.version)
		return
	}

	if uint64(t.version) <= p.index+pow(2, p.layer-1) {
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
