package JsonArrayToObjectConvertInteractorMock

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"

type Interactor struct {
	jsonEntityIn  JsonEntityInterace.JsonEntity
	jsonEntityOut JsonEntityInterace.JsonEntity

	simulateErrorStepMatch int
	simulateErrorFlag      bool
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetJsonEntityIn(jsonEntityIn JsonEntityInterace.JsonEntity) *Interactor {
	i.jsonEntityIn = jsonEntityIn
	return i
}

func (i *Interactor) SetJsonEntityOut(jsonEntityOut JsonEntityInterace.JsonEntity) *Interactor {
	i.jsonEntityOut = jsonEntityOut
	return i
}
