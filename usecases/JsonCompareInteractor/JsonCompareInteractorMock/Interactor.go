package JsonCompareInteractorMock

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"

type Interactor struct {
	first  JsonEntityInterace.JsonEntity
	second JsonEntityInterace.JsonEntity

	isEqual bool

	simulateErrorStepMatch int
	simulateErrorFlag      bool
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetFirst(first JsonEntityInterace.JsonEntity) *Interactor {
	i.first = first
	return i
}

func (i *Interactor) SetSecond(second JsonEntityInterace.JsonEntity) *Interactor {
	i.second = second
	return i
}

func (i *Interactor) SetIsEqual(isEqual bool) *Interactor {
	i.isEqual = isEqual
	return i
}
