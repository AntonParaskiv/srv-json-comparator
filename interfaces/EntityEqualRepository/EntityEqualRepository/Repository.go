package EntityEqualRepository

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/EntityEqualRepository/EquallerInterface"
)

type Repository struct {
	equaller EquallerInterface.Equaller
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetEqualler(equaller EquallerInterface.Equaller) *Repository {
	r.equaller = equaller
	return r
}

func (r *Repository) IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool) {
	isEqual = r.equaller.IsEqual(first, second)
	return
}
