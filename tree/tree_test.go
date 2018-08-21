package tree

import (
	"fmt"
	"testing"
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
		{[]byte{0x0}, []byte{0x0}},
		{[]byte{0x1}, []byte{0x1}},
	}

	for _, test := range events {
		digest := tree.Add(test.event)
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
