package hasher

import (
	"crypto/sha256"
	"hash"
)

type Hashashin struct {
	h hash.Hash
}

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

func NewHashahin() *Hashashin {
	return &Hashashin{h: sha256.New()}
}
