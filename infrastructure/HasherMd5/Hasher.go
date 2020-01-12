package HasherMd5

import (
	"crypto/md5"
	"fmt"
)

type Hasher struct {
}

func New() (h *Hasher) {
	h = new(Hasher)
	return
}

func (h *Hasher) CreateHash(data []byte) (hash string) {
	hashData := md5.Sum(data)
	hash = fmt.Sprintf("%x", hashData)
	return
}
