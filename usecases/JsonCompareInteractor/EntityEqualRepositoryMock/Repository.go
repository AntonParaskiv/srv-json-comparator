package EntityEqualRepositoryMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

type Repository struct {
	first   interface{}
	second  interface{}
	isEqual bool
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetFirst(first interface{}) *Repository {
	r.first = first
	return r
}

func (r *Repository) GetFirst() (first interface{}) {
	first = r.first
	return
}

func (r *Repository) SetSecond(second interface{}) *Repository {
	r.second = second
	return r
}

func (r *Repository) GetSecond() (second interface{}) {
	second = r.second
	return
}

func (r *Repository) SetIsEqual(isEqual bool) *Repository {
	r.isEqual = isEqual
	return r
}

func (r *Repository) GetIsEqual() (isEqual bool) {
	isEqual = r.isEqual
	return
}

func (r *Repository) IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool) {
	r.SetFirst(first)
	r.SetSecond(second)
	isEqual = r.GetIsEqual()
	return
}
