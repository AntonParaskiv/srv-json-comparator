package JsonCompareInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
)

func (i *Interactor) IsEqual(first, second JsonEntity.JsonEntity) (isEqual bool, err error) {
	i.SetFirst(first)
	i.SetSecond(second)
	isEqual = i.isEqual

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}
