package JsonCompareInteractorMock

import "github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"

func (i *Interactor) IsEqual(first, second JsonEntityInterace.JsonEntity) (isEqual bool, err error) {
	i.SetFirst(first)
	i.SetSecond(second)
	isEqual = i.isEqual

	if i.IsSetSimulateError() {
		err = i.Error()
		return
	}

	return
}
