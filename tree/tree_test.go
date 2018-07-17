package tree

import (
	"fmt"
	"testing"

	"github.com/suizman/htree/utils/hashing"
)

func TestAdd(t *testing.T) {

	store := Node{
		hashoff: make(map[Pos]Digest),
	}

	hasher := new(hasher.Sha256Hasher)

	tree := NewTree(
		"barbol",
		0,
		hasher,
		store,
	)

	events := []struct {
		event []byte
	}{
		{[]byte{0x0}},
		{[]byte{0x1}},
		{[]byte{0x2}},
	}

	for i, c := range events {
		i++
		digest := tree.Add(c.event)
		fmt.Printf("New digest created %x store: %v\n", digest, store)
	}
}

func TestNewTree(t *testing.T) {

	store := Node{
		hashoff: make(map[Pos]Digest),
	}
	hasher := new(hasher.Sha256Hasher)
	tree := NewTree(
		"barbol",
		0,
		hasher,
		store,
	)

	fmt.Printf("This is your new tree id: %s, version: %v, position: %x\n", tree.treeId, tree.version, tree.store)
}

func TestGetDepth(t *testing.T) {
	depth := getDepth(3)
	fmt.Printf("Actual depth: %v\n", depth)
}
