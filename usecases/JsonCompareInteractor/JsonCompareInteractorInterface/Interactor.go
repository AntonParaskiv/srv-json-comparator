package JsonCompareInteractorInterface

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"

type Interactor interface {
	IsEqual(first, second JsonEntityInterace.JsonEntity) (isEqual bool, err error)
}
