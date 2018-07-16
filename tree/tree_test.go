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
		index := uint64(i)
		digest := tree.Add(c.event, uInt64AsBytes(index))
		fmt.Printf("New digest created %x store: %v\n", digest, store)
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

func TestTravel(t *testing.T) {

	pos := Pos{
		index: 0,
		layer: 1,
	}

	travel := Travel(1, pos)

	fmt.Println(travel)
	// if travel == false {
	// 	t.Errorf("Error: %v\n", travel)
	// }
}

func TestGetDepth(t *testing.T) {
	depth := getDepth(3)
	fmt.Printf("Actual depth: %v\n", depth)
}
