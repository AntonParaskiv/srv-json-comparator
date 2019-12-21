package JsonArrayToObjectConvertInteractorInterface

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"
)

type Interactor interface {
	EntityArraysToObjects(jsonEntityIn JsonEntityInterace.JsonEntity) (jsonEntityOut JsonEntityInterace.JsonEntity, err error)
}
