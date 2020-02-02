package HasherMock

type Hasher struct {
	data []byte
	hash string
}

func New() (h *Hasher) {
	h = new(Hasher)
	return
}

func (h *Hasher) SetData(data []byte) *Hasher {
	h.data = data
	return h
}

func (h *Hasher) GetData() (data []byte) {
	data = h.data
	return
}

func (h *Hasher) SetHash(hash string) *Hasher {
	h.hash = hash
	return h
}

func (h *Hasher) GetHash() (hash string) {
	hash = h.hash
	return
}

func (h *Hasher) CreateHash(data []byte) (hash string) {
	h.SetData(data)
	hash = h.GetHash()
	return
}
