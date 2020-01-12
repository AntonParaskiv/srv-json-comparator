package HashRepository

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HasherInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/MarshallerInterface"
	"github.com/pkg/errors"
)

type Repository struct {
	jsonMarshaller MarshallerInterface.Marshaller
	hasher         HasherInterface.Hasher
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetJsonMarshaller(jsonMarshaller MarshallerInterface.Marshaller) *Repository {
	r.jsonMarshaller = jsonMarshaller
	return r
}

func (r *Repository) SetHasher(hasher HasherInterface.Hasher) *Repository {
	r.hasher = hasher
	return r
}

func (r *Repository) GetHash(entity JsonEntity.JsonEntity) (hash string, err error) {
	var entityJsoned []byte

	entityJsoned, err = r.jsonMarshaller.MarshalBytes(entity)
	if err != nil {
		err = errors.Errorf("marshal entity failed")
		return
	}
	hash = r.hasher.CreateHash(entityJsoned)
	return
}
