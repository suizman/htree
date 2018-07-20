package hasher

import (
	"bytes"
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
