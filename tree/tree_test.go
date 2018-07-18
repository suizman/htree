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
		0,
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
		fmt.Printf("New digest created %x\n", digest)
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
