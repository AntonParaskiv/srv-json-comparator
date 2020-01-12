package HasherInterface

type Hasher interface {
	CreateHash(data []byte) (hash string)
}
