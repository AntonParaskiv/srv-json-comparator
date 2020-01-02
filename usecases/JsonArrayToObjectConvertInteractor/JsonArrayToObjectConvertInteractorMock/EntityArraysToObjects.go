package JsonArrayToObjectConvertInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

func (i *Interactor) EntityArraysToObjects(jsonEntityIn JsonEntity.JsonEntity) (jsonEntityOut JsonEntity.JsonEntity, err error) {
	i.SetJsonEntityIn(jsonEntityIn)
	jsonEntityOut = i.jsonEntityOut

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}
