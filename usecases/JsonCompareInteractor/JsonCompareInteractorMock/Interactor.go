package JsonCompareInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

type Interactor struct {
	first  JsonEntity.JsonEntity
	second JsonEntity.JsonEntity

	isEqual bool

	simulateErrorStepMatch int
	simulateErrorFlag      bool
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetFirst(first JsonEntity.JsonEntity) *Interactor {
	i.first = first
	return i
}

func (i *Interactor) SetSecond(second JsonEntity.JsonEntity) *Interactor {
	i.second = second
	return i
}

func (i *Interactor) SetIsEqual(isEqual bool) *Interactor {
	i.isEqual = isEqual
	return i
}
