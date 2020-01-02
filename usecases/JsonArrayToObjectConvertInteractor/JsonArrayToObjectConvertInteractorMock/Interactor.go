package JsonArrayToObjectConvertInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

type Interactor struct {
	jsonEntityIn  JsonEntity.JsonEntity
	jsonEntityOut JsonEntity.JsonEntity

	simulateErrorStepMatch int
	simulateErrorFlag      bool
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetJsonEntityIn(jsonEntityIn JsonEntity.JsonEntity) *Interactor {
	i.jsonEntityIn = jsonEntityIn
	return i
}

func (i *Interactor) SetJsonEntityOut(jsonEntityOut JsonEntity.JsonEntity) *Interactor {
	i.jsonEntityOut = jsonEntityOut
	return i
}
