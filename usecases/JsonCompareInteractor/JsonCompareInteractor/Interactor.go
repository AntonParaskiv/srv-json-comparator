package JsonCompareInteractor

import "github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractorInterface"

type Interactor struct {
	arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetArrayToObjectConverter(arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor) *Interactor {
	i.arrayToObjectConverter = arrayToObjectConverter
	return i
}
