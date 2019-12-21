package JsonArrayToObjectConvertInteractorMock

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"

func (i *Interactor) EntityArraysToObjects(jsonEntityIn JsonEntityInterace.JsonEntity) (jsonEntityOut JsonEntityInterace.JsonEntity, err error) {
	i.SetJsonEntityIn(jsonEntityIn)
	jsonEntityOut = i.jsonEntityOut

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}
