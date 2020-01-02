package JsonCompareInteractorInterface

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

type Interactor interface {
	IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool, err error)
}
