package JsonArrayToObjectConvertInteractorInterface

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

type Interactor interface {
	EntityArraysToObjects(jsonEntityIn JsonEntity.JsonEntity) (jsonEntityOut JsonEntity.JsonEntity, err error)
}
