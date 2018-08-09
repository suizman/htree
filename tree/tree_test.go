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
		{[]byte{0x3}},
		{[]byte{0x4}},
	}

	for i, c := range events {
		i++
		digest := tree.Add(c.event)
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

func TestMembershipGen(t *testing.T) {
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
		tree.Add(c.event)
	}

	// expectedV1 := Tree{
	// 	treeId:  []byte("barbol"),
	// 	version: 1,
	// 	hasher:  sha256.New(),
	// 	store: Node{
	// 		hashoff: make(map[Pos]Digest),
	// 	},
	// }

	_, err := tree.MembershipGen(5, 0, 5)

	if err != nil {
		t.Fatalf("Error: %v\n", err)
	}
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
