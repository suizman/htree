package hasher

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hash"
)

type Hasher interface {
	Do(...[]byte) []byte
}

type Sha256Hasher struct {
	underlying hash.Hash
}

func NewSha256Hasher() *Sha256Hasher {
	return &Sha256Hasher{underlying: sha256.New()}
}

func (s Sha256Hasher) Do(data ...[]byte) []byte {
	s.underlying.Reset()

	for i := 0; i < len(data); i++ {
		s.underlying.Write(data[i])
	}
	return s.underlying.Sum(nil)[:]
}

func StringHash(s string) string {

	var b bytes.Buffer

	for i := 0; i < len(s); i++ {
		data := fmt.Sprintf("%#x", s[i:i+2])
		b.WriteString(data)
		i = i + 1
	}

	data := fmt.Sprintf("%v", b.String())
	return data
}
