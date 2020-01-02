package JsonArrayToObjectConvertInteractor

import "github.com/AntonParaskiv/srv-json-comparator/usecases/JsonMarshaller"

type Interactor struct {
	jsonMarshaller JsonMarshaller.JsonMarshaller
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}
