package main

import (
	"crypto/sha256"
	"fmt"
)

type Hasher interface {
	Do(...[]byte) []byte
}

type Sha256Hasher struct{}

func (s Sha256Hasher) Do(data ...[]byte) []byte {
	hasher := sha256.New()

	for i := 0; i < len(data); i++ {
		hasher.Write(data[i])
	}

	return hasher.Sum(nil)[:]
}

func main() {
	hasher := new(Sha256Hasher)

	h1 := hasher.Do([]byte("Hello 1"))
	h2 := hasher.Do([]byte("Hello 2"))
	hofh := hasher.Do([]byte(h1), []byte(h2))
	fmt.Printf("Hash 1: %x\n", h1)
	fmt.Printf("Hash 2: %x\n", h2)
	fmt.Printf("Hofh  : %x\n", hofh)
}
