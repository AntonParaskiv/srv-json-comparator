package HashRepository

import (
	"errors"
	"fmt"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HasherInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/MarshallerInterface"
)

type Repository struct {
	marshaller MarshallerInterface.Marshaller
	hasher     HasherInterface.Hasher
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetMarshaller(marshaller MarshallerInterface.Marshaller) *Repository {
	r.marshaller = marshaller
	return r
}

func (r *Repository) SetHasher(hasher HasherInterface.Hasher) *Repository {
	r.hasher = hasher
	return r
}

func (r *Repository) GetHash(entity JsonEntity.JsonEntity) (hash string, err error) {
	var entityMarshalled []byte

	entityMarshalled, err = r.marshaller.MarshalBytes(entity)
	if err != nil {
		err = errorMarshalEntityFailed(err)
		return
	}
	hash = r.hasher.CreateHash(entityMarshalled)
	return
}

func errorMarshalEntityFailed(wrapErr error) (err error) {
	err = wrapError("marshal entity failed", wrapErr)
	return err
}

func wrapError(msg string, wrapErr error) (err error) {
	if wrapErr == nil {
		err = errors.New(msg)
	} else {
		err = fmt.Errorf("%s: %w", msg, wrapErr)
	}
	return
}
