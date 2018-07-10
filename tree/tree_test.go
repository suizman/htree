package tree

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	position := make(map[Pos]Digest)

	node := Node{
		position,
	}

	tree := NewTree(
		"barbol",
		0,
		node,
	)

	event := []byte("Test event")

	digest := tree.Add(event, []byte("0"))

	fmt.Printf("New digest created %x\n", digest)
}

func TestNewTree(t *testing.T) {

	position := make(map[Pos]Digest)

	node := Node{
		position,
	}

	tree := NewTree(
		"barbol",
		0,
		node,
	)

	fmt.Printf("This is your new tree id: %s, version: %v, position: %x\n", tree.treeId, tree.version, tree.node)
}
