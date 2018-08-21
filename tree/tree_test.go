package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	store := Node{
		hashoff: make(map[Pos]Digest),
	}

	tree := NewTree(
		"barbol",
		-1,
		store,
	)

	events := []struct {
		event          []byte
		expectedDigest []byte
	}{
		{[]byte("0"), []byte{0x5f, 0xec, 0xeb, 0x66, 0xff, 0xc8, 0x6f, 0x38, 0xd9, 0x52, 0x78, 0x6c, 0x6d, 0x69, 0x6c, 0x79, 0xc2, 0xdb, 0xc2, 0x39, 0xdd, 0x4e, 0x91, 0xb4, 0x67, 0x29, 0xd7, 0x3a, 0x27, 0xfb, 0x57, 0xe9}},
		{[]byte("1"), []byte{0xfa, 0x13, 0xbb, 0x36, 0xc0, 0x22, 0xa6, 0x94, 0x3f, 0x37, 0xc6, 0x38, 0x12, 0x6a, 0x2c, 0x88, 0xfc, 0x8d, 0x00, 0x8e, 0xb5, 0xa9, 0xfe, 0x8f, 0xcd, 0xe1, 0x70, 0x26, 0x80, 0x7f, 0xea, 0xe4}},
		//{[]byte("2"), []byte{0x93, 0x8d, 0xb8, 0xc9, 0xf8, 0x2c, 0x8c, 0xb5, 0x8d, 0x3f, 0x3e, 0xf4, 0xfd, 0x25, 0x0, 0x36, 0xa4, 0x8d, 0x26, 0xa7, 0x12, 0x75, 0x3d, 0x2f, 0xde, 0x5a, 0xbd, 0x3, 0xa8, 0x5c, 0xab, 0xf4}},
	}

	for _, test := range events {
		digest := tree.Add(test.event)
		assert.Equalf(t, test.expectedDigest, digest, "Unexpected digest")
		fmt.Printf("New root digest created: %x\n", digest)
	}

}

func TestNewTree(t *testing.T) {

	store := Node{
		hashoff: make(map[Pos]Digest),
	}

	tree := NewTree(
		"barbol",
		0,
		store,
	)

	fmt.Printf("This is your new tree id: %s, version: %v, position: %x\n", tree.treeId, tree.version, tree.store)
}

func TestGenProof(t *testing.T) {
	digest := Digest{value: []byte("digest")}
	store := Node{
		hashoff: make(map[Pos]Digest),
	}

	tree := NewTree(
		"barbol",
		2,
		store,
	)

	store.hashoff[Pos{index: 0, layer: 0}] = digest
	store.hashoff[Pos{index: 0, layer: 1}] = digest
	store.hashoff[Pos{index: 1, layer: 0}] = digest
	store.hashoff[Pos{index: 0, layer: 2}] = digest
	store.hashoff[Pos{index: 2, layer: 1}] = digest
	store.hashoff[Pos{index: 2, layer: 0}] = digest
	store.hashoff[Pos{index: 3, layer: 0}] = digest
	tree.GenProof(2, digest.value)
}

func TestGetDepth(t *testing.T) {
	store := Node{
		hashoff: make(map[Pos]Digest),
	}

	tree := NewTree(
		"barbol",
		0,
		store,
	)

	depth := tree.getDepth()

	fmt.Printf("Actual depth: %v\n", depth)
}

func TestGetVersion(t *testing.T) {
	store := Node{
		hashoff: make(map[Pos]Digest),
	}

	tree := NewTree(
		"barbol",
		0,
		store,
	)

	for i := 0; i < 10; i++ {
		tree.version++
	}

	version := tree.GetVersion()

	fmt.Printf("Actual version: %v\n", version)
}
