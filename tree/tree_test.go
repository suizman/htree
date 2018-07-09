package tree

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	tree := &Tree{[]byte("tree-fake-id")}
	event := []byte("Test event")

	digest := tree.Add(event, []byte("0"))

	fmt.Printf("New digest created %x\n", digest)
}

func TestNewTree(t *testing.T) {

	tree := NewTree("barbol")
	fmt.Printf("This is your new tree %s\n", tree)
}
